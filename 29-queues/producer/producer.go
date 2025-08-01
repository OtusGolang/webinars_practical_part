// Этот пример создает постоянную (durable) точку обмена (Exchange) и публикует одно сообщение
// в эту точку обмена с заданным ключом маршрутизации (routing key).
package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// publish — пример публикации сообщения в Exchange.
// Этот код основан на официальном примере: https://github.com/streadway/amqp/blob/master/_examples/simple-producer/producer.go
func publish(amqpURI, exchange, exchangeType, routingKey, body string, reliable bool) error {
	// Эта функция выполняет подключение, открывает канал, объявляет Exchange, публикует сообщение и закрывает соединение.
	// В реальном сервисе обычно поддерживается постоянное соединение и публикация происходит через него.

	log.Printf("dialing %q", amqpURI)
	connection, err := amqp.Dial(amqpURI)
	if err != nil {
		return fmt.Errorf("dial: %s", err)
	}
	defer connection.Close()

	log.Printf("got Connection, getting Channel")
	channel, err := connection.Channel()
	if err != nil {
		return fmt.Errorf("channel: %s", err)
	}

	log.Printf("got Channel, declaring %q Exchange (%q)", exchangeType, exchange)
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

	// Надежные подтверждения публикации требуют поддержки confirm.select от соединения.
	if reliable {
		log.Printf("enabling publishing confirms.")
		if err := channel.Confirm(false); err != nil {
			return fmt.Errorf("channel could not be put into confirm mode: %s", err)
		}

		confirms := channel.NotifyPublish(make(chan amqp.Confirmation, 1))

		defer confirmOne(confirms)
	}

	log.Printf("declared Exchange, publishing %dB body (%q)", len(body), body)
	if err = channel.Publish(
		exchange,   // publish to an exchange
		routingKey, // routing to 0 or more queues
		false,      // mandatory
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

	return nil
}

// Обычно поддерживается канал публикаций, счетчик последовательности и множество неподтвержденных номеров,
// и выполняется цикл до закрытия канала публикаций.
func confirmOne(confirms <-chan amqp.Confirmation) {
	log.Printf("waiting for confirmation of one publishing")

	if confirmed := <-confirms; confirmed.Ack {
		log.Printf("confirmed delivery with delivery tag: %d", confirmed.DeliveryTag)
	} else {
		log.Printf("failed delivery of delivery tag: %d", confirmed.DeliveryTag)
	}
}
