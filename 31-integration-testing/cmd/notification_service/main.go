package main

import (
	"context"
	"io"
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/streadway/amqp"
)

func main() {
	var conf config
	failOnError(envconfig.Process("notify_service", &conf), "failed to init config")

	// Init RabbitMQ
	conn, err := amqp.Dial(conf.AmqpDSN)
	failOnError(err, "failed to connect to RabbitMQ")
	defer failOnClose(conn, "failed to close RMQ connection")

	ch, err := conn.Channel()
	failOnError(err, "failed to open RMQ channel")
	defer failOnClose(ch, "failed to close RMQ channel")

	s := sender{ch}
	w, err := s.Process(context.TODO(), conf.QueueName, conf.RegExchangeName, conf.NotifyExchangeName)
	failOnError(err, "failed to start process user registrations")

	log.Println("wait for new user registrations...")
	w.Wait()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func failOnClose(closer io.Closer, msg string) func() {
	return func() {
		failOnError(closer.Close(), msg)
	}
}
