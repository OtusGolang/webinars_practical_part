package middleware

import (
	"context"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var requestCnt int64

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	requestId := atomic.AddInt64(&requestCnt, 1)
	ctx := context.WithValue(r.Context(), "request_id", requestId)
	rCtx := r.Clone(ctx)
	l.handler.ServeHTTP(w, rCtx)
	log.Printf("%s %s %v request_id: %d", r.Method, r.URL.Path, time.Since(start), requestId)
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

func IsArgExists(h http.HandlerFunc, argName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		args := r.URL.Query()
		arg := args.Get(argName)
		if len(arg) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		h(w, r)
	}
}
