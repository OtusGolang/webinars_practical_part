package testingm

// ─────────────────────────────────────────────────────────────────────────────
// testing.M и функция TestMain — точка входа всего тестового бинарника пакета.
//
// ЧТО ЭТО ТАКОЕ
// Если в пакете объявлена функция с сигнатурой
//
//	func TestMain(m *testing.M)
//
// то go test передаёт управление ей, а не запускает тесты напрямую.
// Ни один Test*/Benchmark*/Fuzz*/Example* не стартует, пока вы сами
// не вызовете m.Run(). Это единственное место, где можно выполнить код
// ДО первого теста и ПОСЛЕ последнего.
//
// ЧТО ПРОИСХОДИТ ПОД КАПОТОМ
// `go test` — это кодогенератор. Он собирает отдельный main-пакет и кладёт его
// во временный каталог сборки ($GOCACHE, посмотреть можно так:
//
//	go test -work -run TestSetGet ./05-testing/testingm/
//	# в WORK=... лежит b0??/_testmain.go
//
// Сгенерированный _testmain.go выглядит примерно так:
//
//	var tests = []testing.InternalTest{
//	    {"TestSetGet", testingm.TestSetGet},
//	    ...
//	}
//
//	func main() {
//	    m := testing.MainStart(testdeps.TestDeps{}, tests, benchmarks, fuzzTargets, examples)
//	    testingm.TestMain(m)                       // ← вот сюда попадаем мы
//	    os.Exit(int(reflect.ValueOf(m).Elem().
//	        FieldByName("exitCode").Int()))        // ← а вот так Go 1.15+ достаёт код
//	}
//
// Если TestMain НЕ объявлен, генерируется просто `os.Exit(m.Run())`.
//
// Отсюда следуют все практические выводы:
//   - тесты — это обычные функции в слайсе, регистрируются на этапе сборки;
//   - TestMain выполняется в главной горутине, *testing.T здесь нет;
//   - os.Exit(code) в конце не обязателен начиная с Go 1.15: код возврата
//     достанут рефлексией из приватного поля m.exitCode. Раз поле приватное,
//     поменять вердикт ПОСЛЕ m.Run() можно только через os.Exit
//     (см. проверку на утечку горутин ниже);
//   - TestMain в пакете может быть только один: объявление одновременно
//     в `testingm` и в `testingm_test` — ошибка сборки
//     "multiple definitions of TestMain".
// ─────────────────────────────────────────────────────────────────────────────

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"testing"
	"time"
)

// Общий на весь пакет ресурс. Инициализируется один раз в TestMain,
// тесты его только используют (и не имеют права закрывать).
var (
	testStore *Store
	testDir   string
)

// Свои флаги тестового бинарника.
// Объявлять их нужно на уровне пакета (то есть до flag.Parse), а не внутри TestMain.
// Передаются так (ПОСЛЕ списка пакетов — иначе `go test` съест "./testingm/"
// как значение неизвестного ему флага и запустит не тот пакет):
//
//	go test ./testingm/ -keep-data
//	go test ./testingm/ -args -keep-data   # всё после -args уходит в бинарник как есть
var keepData = flag.Bool("keep-data", false, "не удалять временный каталог после тестов")

// TestMain — «main()» тестового бинарника. Современный вариант (Go 1.15+).
func TestMain(m *testing.M) {
	// 1. РАЗБОР ФЛАГОВ.
	// m.Run() сам вызовет flag.Parse(), если этого ещё не сделали.
	// Но нам флаги нужны РАНЬШЕ — в setup(). Плюс testing.Short() и
	// testing.Verbose() до разбора флагов паникуют/врут, поэтому парсим явно.
	flag.Parse()

	// 2. SETUP.
	// Здесь нет *testing.T, значит нет ни t.Fatal, ни t.TempDir, ни t.Cleanup.
	// Об ошибке сообщаем в stderr и выходим ненулевым кодом.
	if err := setup(); err != nil {
		fmt.Fprintf(os.Stderr, "setup failed: %v\n", err)
		// return здесь дал бы «зелёный» прогон: m.Run() не вызывался, m.exitCode == 0.
		// Тестов и отчёта о покрытии тоже не будет.
		os.Exit(1)
	}

	// 3. TEARDOWN ЧЕРЕЗ defer.
	// Главное правило файла: os.Exit не выполняет отложенные вызовы.
	// Поэтому defer работает только там, где мы выходим через return, —
	// на каждом os.Exit ниже teardown придётся звать руками.
	defer teardown()

	// Типичный приём: пропустить весь пакет, если окружения нет.
	// os.Exit(0) — «зелёный» выход без запуска тестов.
	if os.Getenv("SKIP_TESTINGM") != "" {
		fmt.Fprintln(os.Stderr, "SKIP_TESTINGM is set, skipping package")
		teardown()
		os.Exit(0)
	}

	before := runtime.NumGoroutine()

	// 4. ЗАПУСК ТЕСТОВ.
	// m.Run() блокируется до завершения ВСЕХ тестов пакета, включая t.Parallel()
	// и Example-функции. Возвращаемое значение можно игнорировать: ровно тот же код
	// уже лежит в приватном m.exitCode, откуда его и заберёт сгенерированный main().
	// Вызвать m.Run() можно и несколько раз — см. README, раздел «Матрица конфигураций».
	m.Run()

	// 5. ПОСТПРОВЕРКИ ВСЕГО ПАКЕТА.
	// После m.Run() можно проверить то, что не видно изнутри отдельного теста:
	// утечку горутин, незакрытые соединения, обращения к мокам.
	// Так работает go.uber.org/goleak: goleak.VerifyTestMain(m).
	if leaked := runtime.NumGoroutine() - before; leaked > 0 {
		// Даём планировщику добить уже завершающиеся горутины, иначе будут ложные срабатывания.
		time.Sleep(50 * time.Millisecond)
		if leaked = runtime.NumGoroutine() - before; leaked > 0 {
			fmt.Fprintf(os.Stderr, "goroutine leak detected: +%d\n", leaked)
			// ВАЖНО: «ужесточить» вердикт после m.Run() можно ТОЛЬКО через os.Exit.
			// m.exitCode приватный, снаружи его не поменять, а return вернёт
			// тот код, который записал m.Run() (у нас — 0, тесты-то прошли).
			teardown()
			os.Exit(1)
		}
	}

	// 6. ВЫХОД.
	// Просто return: код возврата сгенерированный main() достанет из m.exitCode.
}

// Как то же самое писали до Go 1.15 (и как до сих пор пишут по привычке —
// этот вариант корректен в любой версии Go, просто более шумный):
//
//	func TestMain(m *testing.M) {
//		flag.Parse()
//
//		if err := setup(); err != nil {
//			fmt.Fprintf(os.Stderr, "setup failed: %v\n", err)
//			os.Exit(1)
//		}
//
//		code := m.Run()
//
//		teardown()    // только руками, defer тут бесполезен
//		os.Exit(code) // обязательно, иначе прогон всегда «зелёный»
//	}
//
// Компромисс, если нужны и defer, и свой код возврата (например, как у нас —
// уронить прогон из-за утечки горутин): вынести тело в отдельную функцию.
//
//	func TestMain(m *testing.M) { os.Exit(run(m)) }
//
//	func run(m *testing.M) int {
//		flag.Parse()
//		if err := setup(); err != nil { ... return 1 }
//		defer teardown()   // отработает: os.Exit зовётся уже снаружи
//		return m.Run()
//	}

func setup() error {
	var err error

	// Аналог t.TempDir(), только без автоочистки — убираем сами в teardown.
	testDir, err = os.MkdirTemp("", "testingm-*")
	if err != nil {
		return fmt.Errorf("create temp dir: %w", err)
	}

	testStore, err = Open(testDir)
	if err != nil {
		return fmt.Errorf("open store: %w", err)
	}

	// testing.Verbose() доступен только после flag.Parse().
	if testing.Verbose() {
		fmt.Fprintf(os.Stderr, "[setup] store dir: %s\n", testDir)
	}

	return nil
}

func teardown() {
	if testStore != nil {
		if err := testStore.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "[teardown] close store: %v\n", err)
		}
	}

	if *keepData {
		fmt.Fprintf(os.Stderr, "[teardown] -keep-data: данные оставлены в %s\n", testDir)
		return
	}

	if err := os.RemoveAll(testDir); err != nil {
		fmt.Fprintf(os.Stderr, "[teardown] remove temp dir: %v\n", err)
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// Дальше — обычные тесты. Они ничего не знают про setup/teardown и просто
// пользуются общим ресурсом. Правило: тест НЕ закрывает и НЕ пересоздаёт
// общий ресурс, а изолируется через уникальные ключи/схемы/транзакции.
// ─────────────────────────────────────────────────────────────────────────────

func TestSetGet(t *testing.T) {
	const key, want = "TestSetGet/user", "gopher"

	if err := testStore.Set(key, want); err != nil {
		t.Fatalf("Set(%q, %q) returns unexpected error: %v", key, want, err)
	}

	got, ok := testStore.Get(key)
	if !ok {
		t.Fatalf("Get(%q): key not found", key)
	}
	if got != want {
		t.Errorf("Get(%q) = %q; want %q", key, got, want)
	}
}

func TestGetMissingKey(t *testing.T) {
	if _, ok := testStore.Get("TestGetMissingKey/nope"); ok {
		t.Error("Get() returns ok for a missing key")
	}
}

// Параллельные тесты тоже завершатся до возврата из m.Run() —
// teardown в TestMain гарантированно не выдернет ресурс из-под них.
func TestSetParallel(t *testing.T) {
	for _, name := range []string{"alpha", "beta", "gamma"} {
		// go 1.22+: переменная цикла своя на каждой итерации, `name := name` больше не нужен.
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			key := "TestSetParallel/" + name
			if err := testStore.Set(key, name); err != nil {
				t.Fatalf("Set(%q, %q) returns unexpected error: %v", key, name, err)
			}
		})
	}
}

// Пример «долгого» теста: -short обрабатываем сами, testing его не пропускает.
func TestPersistence(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping disk round-trip in -short mode")
	}

	const key, want = "TestPersistence/key", "value"
	if err := testStore.Set(key, want); err != nil {
		t.Fatalf("Set(%q, %q) returns unexpected error: %v", key, want, err)
	}

	// Открываем тот же каталог вторым экземпляром — данные должны быть на диске.
	reopened, err := Open(testDir)
	if err != nil {
		t.Fatalf("Open(%q) returns unexpected error: %v", testDir, err)
	}
	defer reopened.Close()

	got, ok := reopened.Get(key)
	if !ok {
		t.Fatalf("Get(%q) after reopen: key not found", key)
	}
	if got != want {
		t.Errorf("Get(%q) after reopen = %q; want %q", key, got, want)
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// Кейс: тестируем код, который завершает процесс (os.Exit / log.Fatal).
// Приём — перезапуск того же тестового бинарника подпроцессом.
// TestMain здесь важен тем, что даёт стабильную точку входа и env-переключатель.
// ─────────────────────────────────────────────────────────────────────────────

func TestExitCodeInSubprocess(t *testing.T) {
	// Ветка, исполняемая уже ВНУТРИ подпроцесса.
	// Обратите внимание: в подпроцессе снова отработает TestMain (со своим setup),
	// а вот teardown — нет, потому что мы убиваем процесс через os.Exit.
	if os.Getenv("TESTINGM_SUBPROCESS") == "1" {
		fmt.Println("работаю в подпроцессе")
		// os.Exit прямо из теста обрывает процесс мимо TestMain,
		// поэтому временный каталог подпроцесса убираем за собой сами.
		teardown()
		os.Exit(3)
	}

	// Ветка родителя: перезапускаем сами себя (os.Args[0] — это тестовый бинарник)
	// с фильтром -test.run и env-переключателем.
	cmd := exec.Command(os.Args[0], "-test.run=^TestExitCodeInSubprocess$")
	cmd.Env = append(os.Environ(), "TESTINGM_SUBPROCESS=1")

	out, err := cmd.CombinedOutput()

	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) {
		t.Fatalf("subprocess: got err = %v; want *exec.ExitError\noutput:\n%s", err, out)
	}
	if got, want := exitErr.ExitCode(), 3; got != want {
		t.Errorf("subprocess exit code = %d; want %d\noutput:\n%s", got, want, out)
	}
}
