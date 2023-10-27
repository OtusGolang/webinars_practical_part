package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Адрес сервера Redis
		DB:   0,                // Номер базы данных (по умолчанию 0)
	})

	message := "Сообщение для обработки"
	queueName := "myqueue"

	// Отправка сообщения в очередь
	err := client.LPush(ctx, queueName, message).Err()
	if err != nil {
		panic(err)
	}

	// Закрытие клиента Redis
	defer client.Close()
}
