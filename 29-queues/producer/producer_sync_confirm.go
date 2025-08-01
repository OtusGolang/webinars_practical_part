// Этот пример создает постоянную (durable) точку обмена (Exchange) и публикует одно сообщение
// в эту точку обмена с заданным ключом маршрутизации (routing key).
package main

import (
	"fmt"
	"time"

	"log/slog"

	"github.com/streadway/amqp"
)

// publish_sync_confirm — пример публикации сообщения с синхронными подтверждениями.
// Основной недостаток — нельзя публиковать более одного сообщения за раз.
func publish_sync_confirm(amqpURI, exchange, exchangeType, routingKey, body string, reliable bool) error {
	// Эта функция выполняет подключение, открывает канал, объявляет Exchange, публикует сообщение и закрывает соединение.
	// В реальном сервисе обычно поддерживается постоянное соединение и публикация происходит через него.

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
	// Надежные подтверждения публикации требуют поддержки confirm.select от соединения.
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

	// После публикации

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
