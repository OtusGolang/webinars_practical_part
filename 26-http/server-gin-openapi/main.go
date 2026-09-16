package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"slices"
	"syscall"
	"time"

	"github.com/OtusGolang/webinars_practical_part/26-http/handler"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
	"github.com/lmittmann/tint"
)

// Тот же сервис голосования, что и в server-mux, но маршруты описаны через huma:
// схема OpenAPI выводится из типов Go, а не пишется руками и не собирается из комментариев.
//
// Заголовок Content-Type обязателен: huma сверяет его со спекой и на голый
// curl -d (тот шлёт application/x-www-form-urlencoded) ответит 415.
// curl -H "Content-Type: application/json" -d '{"candidate_id": 1, "passport": "4509 123456"}' -X POST 0.0.0.0:8080/vote
// curl 0.0.0.0:8080/stat
// curl "0.0.0.0:8080/stat?sort=id"
// curl 0.0.0.0:8080/stat/1
//
// powershell:
//  curl -uri http://localhost:8080/vote -method post -contenttype 'application/json' -body '{"passport":"a", "candidate_id":123}'
//
// Документация и спека:
//  http://localhost:8080/docs         — UI (Stoplight Elements), huma отдаёт его сам
//  http://localhost:8080/openapi.yaml — спека OpenAPI 3.1
//  http://localhost:8080/openapi.json
//
// Проверка валидации (huma ответит 422 с телом RFC 7807, хендлер не вызовется):
// curl -H "Content-Type: application/json" -d '{"candidate_id": 0, "passport": ""}' -X POST 0.0.0.0:8080/vote
// curl "0.0.0.0:8080/stat?sort=nope"

// --- DTO ---
//
// Из Go выводится автоматически: имена полей (json-тег), типы и форматы
// (uint32 → integer/int32 с minimum: 0, time.Time → string/date-time),
// обязательность (поле без указателя и без omitempty → required).
//
// Руками (тегами) задаётся то, чего в системе типов Go нет:
// doc/example — описания и примеры, minimum/maxLength/pattern — ограничения,
// enum — перечисления (в Go нет enum, вывести его неоткуда).

type CandidateStat struct {
	CandidateID uint32 `json:"candidate_id" doc:"Идентификатор кандидата"`
	Votes       uint32 `json:"votes"        doc:"Количество голосов"`
}

type VoteInput struct {
	Body struct {
		Passport    string `json:"passport"       minLength:"1" maxLength:"64" example:"4509 123456" doc:"Номер паспорта голосующего"`
		CandidateID uint32 `json:"candidate_id"   minimum:"1"                 example:"1"           doc:"Идентификатор кандидата"`
		Note        string `json:"note,omitempty" maxLength:"256"                                   doc:"Произвольный комментарий"`
	}
}

type VoteOutput struct {
	Body struct {
		CandidateID uint32    `json:"candidate_id" doc:"Кандидат, за которого учтён голос"`
		Votes       uint32    `json:"votes"        doc:"Количество голосов после учёта"`
		Time        time.Time `json:"time"         doc:"Время учёта голоса"`
	}
}

type StatsInput struct {
	// Тег path/query/header/cookie говорит huma, откуда брать параметр,
	// и одновременно попадает в спеку как parameter.in.
	Sort string `query:"sort" enum:"votes,id" default:"votes" doc:"Порядок сортировки записей"`
}

type StatsOutput struct {
	Body struct {
		// Исходный StatResponse хранит map[uint32]uint32. В JSON Schema ключи объекта
		// всегда строки, поэтому в контракте удобнее массив записей: и в спеке честно,
		// и клиентам генерируются нормальные типы.
		Records []CandidateStat `json:"records" doc:"Статистика по кандидатам"`
		Total   uint32          `json:"total"   doc:"Суммарное количество голосов"`
		Time    time.Time       `json:"time"    doc:"Время формирования ответа"`
	}
}

type StatByCandidateInput struct {
	CandidateID uint32 `path:"candidate_id" minimum:"1" example:"1" doc:"Идентификатор кандидата"`
}

type StatByCandidateOutput struct {
	Body struct {
		CandidateID uint32    `json:"candidate_id" doc:"Идентификатор кандидата"`
		Votes       uint32    `json:"votes"        doc:"Количество голосов"`
		Time        time.Time `json:"time"         doc:"Время формирования ответа"`
	}
}

// --- Маршруты ---

func registerRoutes(api huma.API, svc *handler.Service) {
	// huma.Operation — это ровно та часть спеки, которую из кода вывести нельзя:
	// метод, путь, человекочитаемые описания, теги, коды ошибок, требования безопасности.
	huma.Register(api, huma.Operation{
		OperationID:   "submit-vote",
		Method:        http.MethodPost,
		Path:          "/vote",
		Summary:       "Отдать голос",
		Description:   "Учитывает голос за кандидата и возвращает актуальный счётчик.",
		Tags:          []string{"voting"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, in *VoteInput) (*VoteOutput, error) {
		// Сюда мы попадаем, только если тело запроса прошло валидацию по схеме:
		// разбор JSON, проверку required и ограничений huma делает до вызова хендлера.
		slog.Info("new vote receive",
			"passport", in.Body.Passport,
			"candidate_id", in.Body.CandidateID,
		)

		svc.Lock()
		svc.Stats[in.Body.CandidateID]++
		votes := svc.Stats[in.Body.CandidateID]
		svc.Unlock()

		out := &VoteOutput{}
		out.Body.CandidateID = in.Body.CandidateID
		out.Body.Votes = votes
		out.Body.Time = time.Now()
		return out, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-stats",
		Method:      http.MethodGet,
		Path:        "/stat",
		Summary:     "Общая статистика",
		Description: "Возвращает голоса по всем кандидатам.",
		Tags:        []string{"stat"},
	}, func(ctx context.Context, in *StatsInput) (*StatsOutput, error) {
		svc.RLock()
		records := make([]CandidateStat, 0, len(svc.Stats))
		var total uint32
		for id, votes := range svc.Stats {
			records = append(records, CandidateStat{CandidateID: id, Votes: votes})
			total += votes
		}
		svc.RUnlock()

		slices.SortFunc(records, func(a, b CandidateStat) int {
			if in.Sort == "id" {
				return int(a.CandidateID) - int(b.CandidateID)
			}
			return int(b.Votes) - int(a.Votes)
		})

		out := &StatsOutput{}
		out.Body.Records = records
		out.Body.Total = total
		out.Body.Time = time.Now()
		return out, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-candidate-stat",
		Method:      http.MethodGet,
		Path:        "/stat/{candidate_id}",
		Summary:     "Статистика по кандидату",
		Tags:        []string{"stat"},
		// Коды ошибок в спеку тоже добавляются вручную: из сигнатуры
		// func(...) (*Out, error) видно только то, что ошибка возможна.
		Errors: []int{http.StatusNotFound},
	}, func(ctx context.Context, in *StatByCandidateInput) (*StatByCandidateOutput, error) {
		svc.RLock()
		votes, ok := svc.Stats[in.CandidateID]
		svc.RUnlock()

		slog.Info("candidate lookup", "candidate_id", in.CandidateID, "found", ok)
		if !ok {
			// huma сам сериализует это в RFC 7807 (application/problem+json)
			// с тем кодом, который зашит в конструктор ошибки.
			return nil, huma.Error404NotFound("candidate not found")
		}

		out := &StatByCandidateOutput{}
		out.Body.CandidateID = in.CandidateID
		out.Body.Votes = votes
		out.Body.Time = time.Now()
		return out, nil
	})
}

// requestLogger — обычная gin-мидлварь: huma не заменяет роутер,
// а надстраивается над ним, поэтому весь стек gin остаётся доступен.
func requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		slog.Info("request",
			"method", c.Request.Method,
			"url", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"duration", time.Since(start),
		)
	}
}

func main() {
	slog.SetDefault(slog.New(tint.NewHandler(os.Stdout, nil)))

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery(), requestLogger())

	// Метаданные спеки (info, servers, схемы авторизации) из кода не выводятся ниоткуда —
	// это ручная часть контракта.
	config := huma.DefaultConfig("Voting API", "1.0.0")
	config.Info.Description = "Пример сервиса голосования с автогенерацией OpenAPI из типов Go."
	config.Servers = []*huma.Server{{URL: "http://localhost:8080", Description: "local"}}
	config.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		// Схема объявлена, но ни одна операция её не требует.
		// Чтобы включить: huma.Operation{Security: []map[string][]string{{"apiKey": {}}}}.
		"apiKey": {Type: "apiKey", In: "header", Name: "X-API-Key"},
	}

	// humagin — адаптер: huma умеет работать поверх gin, echo, chi, fiber
	// и стандартного http.ServeMux, сам оставаясь от роутера независимым.
	api := humagin.New(router, config)

	svc := handler.NewService()
	registerRoutes(api, svc)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		slog.Info("server start on", "addr", server.Addr, "docs", "http://localhost:8080/docs")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("listen failed", "err", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	slog.Info("shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("server shutdown failed", "err", err)
		return
	}
	slog.Info("server stopped gracefully")
}
