package main

import (
	"flag"
	"log"
)

// CLI example:
// go run .\... -sync -key none

var (
	uri          = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	exchangeName = flag.String("exchange", "test-exchange", "Durable AMQP exchange name")
	exchangeType = flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
	routingKey   = flag.String("key", "test-key", "AMQP routing key")
	body         = flag.String("body", "foobar", "Body of message")
	reliable     = flag.Bool("reliable", true, "Wait for the publisher confirmation before exiting")
	sync         = flag.Bool("sync", true, "Use synchronous confirms example")
)

func main() {
	flag.Parse()

	if *sync {
		err := publish_sync_confirm(*uri, *exchangeName, *exchangeType, *routingKey, *body, *reliable)
		if err != nil {
			log.Fatalf("%s", err)
		}
	} else {
		err := publish(*uri, *exchangeName, *exchangeType, *routingKey, *body, *reliable)
		if err != nil {
			log.Fatalf("%s", err)
		}
	}
	log.Printf("published %dB OK", len(*body))
}
