package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

type user struct {
	FirstName string
	Email     string
	Age       uint8
}

type sender struct {
	ch *amqp.Channel
}

func (s *sender) Process(_ context.Context, queueName, regExchangeName, notifyExchangeName string) (*sync.WaitGroup, error) {
	if _, err := s.ch.QueueDeclare(queueName, true, true, true, false, nil); err != nil {
		return nil, fmt.Errorf("queue declare err: %w", err)
	}

	if err := s.ch.QueueBind(queueName, "", regExchangeName, false, nil); err != nil {
		return nil, fmt.Errorf("bind notify queue to exchange err: %w", err)
	}

	users, err := s.ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		return nil, fmt.Errorf("start consuming err: %w", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for userData := range users {
			if userData.ContentType == "plain/text" && string(userData.Body) == "HealthCheck" {
				log.Println("health check")
				continue
			}

			if userData.ContentType != "application/json" {
				log.Printf("unexpected msg content type: %s", userData.ContentType)
				continue
			}

			var user user
			err := json.Unmarshal(userData.Body, &user)
			if err != nil {
				log.Printf("invalid user %s: %v", userData.Body, user)
				continue
			}

			// Эмулируем работу по отправке уведомлений пользователю
			time.Sleep(2 * time.Second)

			err = s.ch.Publish(notifyExchangeName, "email", false, false, amqp.Publishing{
				ContentType: "plain/text",
				Body:        []byte(user.Email),
			})
			if err != nil {
				log.Printf("send notification to %s err: %v", user.Email, err)
			}
		}
	}()
	return wg, nil
}
