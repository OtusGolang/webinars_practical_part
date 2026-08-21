# `testing.M` / `TestMain` — точка входа тестового бинарника

```go
func TestMain(m *testing.M)
```

Одна функция на пакет. Если она объявлена, `go test` вызывает её вместо того,
чтобы сразу запускать тесты. Ни один `Test*`, `Benchmark*`, `Fuzz*`, `Example*`
не стартует, пока вы не вызовете `m.Run()`.

Рабочий пример: [example_test.go](example_test.go), тестируемый код: [store.go](store.go).

```
go test -v ./testingm/
go test -v ./testingm/ -keep-data        # свой флаг — ПОСЛЕ списка пакетов
go test -short ./testingm/
SKIP_TESTINGM=1 go test ./testingm/      # пакет пропущен целиком
go test -work -run TestSetGet ./testingm/ # покажет WORK=... с _testmain.go
```

## Под капотом

`go test` — кодогенератор. Он собирает **отдельный `main`-пакет** и кладёт его во
временный каталог сборки. Файл называется `_testmain.go` и выглядит примерно так:

```go
package main

import (
    "os"
    "reflect"
    "testing"
    "testing/internal/testdeps"

    testingm "github.com/OtusGolang/webinars_practical_part/05-testing/testingm"
)

var tests = []testing.InternalTest{
    {"TestSetGet", testingm.TestSetGet},
    {"TestGetMissingKey", testingm.TestGetMissingKey},
    // ...
}

func main() {
    m := testing.MainStart(testdeps.TestDeps{}, tests, benchmarks, fuzzTargets, examples)

    testingm.TestMain(m)                                          // ← ваш TestMain
    os.Exit(int(reflect.ValueOf(m).Elem().FieldByName("exitCode").Int()))
}
```

Без `TestMain` последние две строки заменяются на `os.Exit(m.Run())`.

Что из этого следует:

| Наблюдение | Следствие |
|---|---|
| Тесты — элементы слайса `[]testing.InternalTest`, собранного на этапе сборки | Регистрация статическая, «динамически добавить тест» нельзя; порядок = порядок файлов (алфавит) и позиций в них — полагаться на него нельзя |
| `TestMain` вызывается из `main()` в главной горутине | Здесь нет `*testing.T`: нет `t.Fatal`, `t.TempDir`, `t.Cleanup`, `t.Helper`. Ошибки — в `stderr` + ненулевой код возврата |
| Код возврата достаётся рефлексией из приватного `m.exitCode` | С Go 1.15 `os.Exit` в конце не обязателен: просто `return`, и тогда отработают `defer`. Обратная сторона: поменять вердикт **после** `m.Run()` можно только через `os.Exit` — а он отменяет `defer` |
| `func (m *M) Run() (code int)` начинается с `defer func() { code = m.exitCode }()` | Возвращаемое значение — копия того же поля, которое прочитает сгенерированный `main()`. В комментарии рядом Go-команда честно пишет, что повторные вызовы `m.Run()` не задумывались, но «в дикой природе такие тесты есть» ([#23129](https://go.dev/issue/23129)) |
| Пакеты инициализируются до `main()` | Глобальные `var` и `init()` тестируемого пакета отработают ДО `TestMain`. Ленивую инициализацию лучше не прятать в `init()` |
| `testdeps.TestDeps` — мост между `testing` и `cmd/go` | Через него работают `-cover`, профили и **кэш результатов**: `go test` записывает, какие файлы и переменные окружения читал тест, и инвалидирует кэш при их изменении |
| `flag.Parse()` вызывает `m.Run()`, если его ещё не звали | `testing.Short()` / `testing.Verbose()` до разбора флагов использовать нельзя — парсите явно в начале `TestMain` |
| Каждый пакет — **свой процесс** | `go test ./...` вызовет `TestMain` столько раз, сколько пакетов. «Один контейнер на весь прогон» так не сделать (нужен внешний оркестратор или общий пакет-хелпер с `sync.Once`) |

## Скелет, который стоит помнить

Современный вариант (Go 1.15+):

```go
func TestMain(m *testing.M) {
    flag.Parse()               // 1. флаги — до setup и до testing.Short()

    if err := setup(); err != nil {
        fmt.Fprintf(os.Stderr, "setup: %v\n", err)
        os.Exit(1)             //    здесь os.Exit обязателен: m.Run не вызывался,
    }                          //    m.exitCode == 0, return дал бы «зелёный» прогон

    defer teardown()           // 2. работает, потому что дальше нет os.Exit

    m.Run()                    // 3. блокируется до конца ВСЕХ тестов, включая t.Parallel;
                               //    результат уже записан в m.exitCode
}                              // 4. просто return
```

До Go 1.15 (и по привычке до сих пор) писали так — вариант корректен в любой версии:

```go
func TestMain(m *testing.M) {
    flag.Parse()
    if err := setup(); err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(1) }

    code := m.Run()

    teardown()                 // только руками, defer тут бесполезен
    os.Exit(code)              // обязательно, иначе прогон всегда «зелёный»
}
```

Если нужны **и** `defer`, **и** свой код возврата (уронить прогон по итогам
постпроверки — например, из-за утечки горутин), — выносим тело в функцию:

```go
func TestMain(m *testing.M) { os.Exit(run(m)) }

func run(m *testing.M) int {
    flag.Parse()
    if err := setup(); err != nil { fmt.Fprintln(os.Stderr, err); return 1 }
    defer teardown()           // отработает: os.Exit зовётся уже снаружи
    return m.Run()
}
```

Почему это вообще вопрос: `m.exitCode` — приватное поле, снаружи его не изменить.
`m.Run()` возвращает копию, и после `m.Run()` «ужесточить» вердикт можно
только вызовом `os.Exit` — а он отменяет все `defer`.

## Грабли

1. **`defer teardown()` перед `os.Exit(m.Run())`** — молчаливо не сработает.
   `os.Exit` не выполняет отложенные вызовы. Или `defer` без `os.Exit`, или `os.Exit`
   с ручным teardown — но не оба сразу.
2. **`os.Exit(0)` / `return` до `m.Run()`** (например, после неудачного setup) —
   «все тесты прошли», хотя не запускался ни один. После setup-ошибки — только `os.Exit(1)`.
3. **`code := m.Run()` … `os.Exit(0)`** или потерянный `code` в старом стиле —
   упавшие тесты дадут зелёный CI.
4. **Тест закрывает общий ресурс** (`defer testStore.Close()`) — ломает соседние тесты,
   особенно параллельные. Общий ресурс закрывает только `teardown`.
5. **Два `TestMain`** — в `pkg` и в `pkg_test` — ошибка сборки
   `multiple definitions of TestMain`.
6. **Тяжёлый setup в пакете, где он не нужен** — он выполняется даже при
   `go test -run TestNothingMatches`, потому что `TestMain` вызывается всегда.
7. **`t.Parallel()` + общее изменяемое состояние** — изолируйтесь уникальными
   ключами / схемами БД / транзакциями, а не общим `sync.Mutex` в тестах.

## Типовые use cases

### 1. Дорогой внешний ресурс на весь пакет
БД, docker-контейнер (`testcontainers-go`), брокер, прогретый HTTP-сервер, миграции.
Это основной сценарий — он и реализован в примере (там роль БД играет каталог на диске).

### 2. Пропустить пакет целиком, если окружения нет

```go
if _, err := exec.LookPath("docker"); err != nil {
    fmt.Fprintln(os.Stderr, "docker not found, skipping integration package")
    os.Exit(0)   // зелёный выход без запуска тестов
}
```

### 3. Свои флаги и `-update` для golden-файлов

```go
var update = flag.Bool("update", false, "update .golden files")
```

См. соседний пример [../golden_files/golden_test.go](../golden_files/golden_test.go).

### 4. Постпроверки всего пакета: утечка горутин

```go
import "go.uber.org/goleak"

func TestMain(m *testing.M) {
    goleak.VerifyTestMain(m)   // сам зовёт m.Run() и os.Exit
}
```

В примере то же самое сделано вручную через `runtime.NumGoroutine()`.

### 5. Глобальная настройка окружения
Подмена глобального логгера на «тихий», фиксированная `TZ`, `os.Chdir` в корень
репозитория, детерминированный seed, отключение внешних HTTP-вызовов
(`http.DefaultTransport` → мок).

### 6. Матрица конфигураций: `m.Run()` несколько раз

```go
func TestMain(m *testing.M) {
    flag.Parse()

    var code int
    for _, backend := range []string{"memory", "postgres"} {
        currentBackend = backend       // глобальная настройка, которую читают тесты
        fmt.Fprintf(os.Stderr, "=== backend: %s\n", backend)
        if c := m.Run(); c != 0 {
            code = c
        }
    }

    os.Exit(code)
}
```

Работает, но имена тестов в выводе повторяются — для CI-парсеров это неудобно;
чаще предпочтительнее подтесты `t.Run(backend, ...)`.

### 7. Тест кода, который завершает процесс (`os.Exit`, `log.Fatal`)
Перезапуск того же бинарника подпроцессом: `exec.Command(os.Args[0], "-test.run=^TestX$")`
плюс env-переключатель. Реализовано в `TestExitCodeInSubprocess`.
Так тестируется, например, `os.Exit` в `cmd/`-пакетах и поведение при панике.

### 8. Сбор общих метрик прогона
Замер общего времени, количество запросов к моку, дамп профиля после `m.Run()` —
всё, что нужно посчитать «по пакету», а не «по тесту».
