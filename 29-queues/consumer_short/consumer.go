package simpleconsumer

import (
	"context"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RMQConnection interface {
	Channel() (*amqp.Channel, error)
}

type Consumer struct {
	name string
	conn RMQConnection
}

func New(name string, conn RMQConnection) *Consumer {
	return &Consumer{
		name: name,
		conn: conn,
	}
}

type Message struct {
	Ctx  context.Context
	Data []byte
}

func (c *Consumer) Consume(ctx context.Context, queue string) (<-chan Message, error) {
	messages := make(chan Message)

	ch, err := c.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("open channel: %w", err)
	}

	go func() {
		<-ctx.Done()
		if err := ch.Close(); err != nil {
			log.Println(err)
		}
	}()

	deliveries, err := ch.Consume(queue, c.name, false, false, false, false, nil)
	if err != nil {
		return nil, fmt.Errorf("start consuming: %w", err)
	}

	go func() {
		defer func() {
			close(messages)
			log.Println("close messages channel")
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case del := <-deliveries:
				if err := del.Ack(false); err != nil {
					log.Println(err)
				}

				msg := Message{
					Ctx:  context.TODO(),
					Data: del.Body,
				}

				select {
				case <-ctx.Done():
					return
				case messages <- msg:
				}
			}
		}
	}()

	return messages, nil
}
