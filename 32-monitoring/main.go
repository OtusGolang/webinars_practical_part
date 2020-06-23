package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	middlewarestd "github.com/slok/go-http-metrics/middleware/std"
)

func myHandler(w http.ResponseWriter, _ *http.Request) {
	defer regCounter.Inc()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world!"))
}

func main() {
	// Middleware для мониторинга
	mdlw := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})
	h := middlewarestd.Handler("", mdlw, http.HandlerFunc(myHandler))

	// HTTP exporter для prometheus
	go http.ListenAndServe(":9091", promhttp.Handler())

	// Ваш основной HTTP сервис
	if err := http.ListenAndServe(":9092", h); err != nil {
		log.Panicf("error while serving: %s", err)
	}
}
