package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var regCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "business_registration",
	Help: "Client registration event",
})

func init() {
	prometheus.MustRegister(regCounter)
}
