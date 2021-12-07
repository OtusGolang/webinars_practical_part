package middleware

import (
	"log"
	"net/http"
	"time"
)

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
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
