// This example declares a durable Exchange, and publishes a single message to
// that Exchange with a given routing key.
package main

import (
	"fmt"
	"time"

	"log/slog"

	"github.com/streadway/amqp"
)

// publish_sync_confirm is an example of publishing a message with synchronous confirms.
// Main disadvadage here is thet we canot publesh more than one message at a time.
func publish_sync_confirm(amqpURI, exchange, exchangeType, routingKey, body string, reliable bool) error {
	// This function dials, connects, declares, publishes, and tears down,
	// all in one go. In a real service, you probably want to maintain a
	// long-lived connection as state, and publish against that.

	slog.Info("dialing", "amqpURI", amqpURI)
	connection, err := amqp.Dial(amqpURI)
	if err != nil {
		return fmt.Errorf("dial: %s", err)
	}
	defer connection.Close()

	slog.Info("got Connection, getting Channel")
	channel, err := connection.Channel()
	if err != nil {
		return fmt.Errorf("channel: %s", err)
	}

	slog.Info("got Channel, declaring Exchange", "exchangeType", exchangeType, "exchange", exchange)
	if err := channel.ExchangeDeclare(
		exchange,     // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return fmt.Errorf("exchange Declare: %s", err)
	}
	confirms := channel.NotifyPublish(make(chan amqp.Confirmation, 1))
	// Reliable publisher confirms require confirm.select support from the
	// connection.
	if reliable {
		slog.Info("enabling publishing confirms.")
		if err := channel.Confirm(false); err != nil {
			return fmt.Errorf("channel could not be put into confirm mode: %s", err)
		}

		//defer confirmOne(confirms)
	}

	slog.Info("declared Exchange, publishing body", "bytes", len(body), "body", body)

	returns := channel.NotifyReturn(make(chan amqp.Return, 1))

	if err = channel.Publish(
		exchange,   // publish to an exchange
		routingKey, // routing to 0 or more queues
		true,       // mandatory: true to get unrouted messages returned
		false,      // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(body),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		return fmt.Errorf("exchange Publish: %s", err)
	}

	// After publishing

	select {
	case ret := <-returns:
		slog.Warn("Message was returned", "body", string(ret.Body))
		slog.Warn("message was delivered to exchange but not routed to any queue")
	case conf := <-confirms:
		if conf.Ack {
			slog.Info("Message confirmed by broker")
		} else {
			return fmt.Errorf("message was not confirmed by broker")
		}
	case <-time.After(1 * time.Second):
		return fmt.Errorf("timeout waiting for broker response")
	}

	return nil
}
