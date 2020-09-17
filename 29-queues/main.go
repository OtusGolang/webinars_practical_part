package main

import (
	"context"
	"flag"
	simpleconsumer "github.com/OtusGolang/webinars_practical_part/29-queues/consumer"
	"github.com/streadway/amqp"
	"log"
)

var (
	uri = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
)

func init() {
	flag.Parse()
}

// http://localhost:15672/ guest:guest
func main() {
	conn, err := amqp.Dial(*uri)
	failOnErr(err)

	c := simpleconsumer.New("simple consumer", conn)
	msgs, err := c.Consume(context.Background(), "hello")

	log.Println("start consuming...")

	for m := range msgs {
		log.Println("receive new message: ", string(m.Data))
	}
}

func failOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
