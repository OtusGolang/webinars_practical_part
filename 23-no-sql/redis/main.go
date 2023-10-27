package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	// Создание клиента Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Адрес и порт Redis-сервера
		Password: "",               // Пароль (если есть)
		DB:       0,                // Номер базы данных
	})

	// Проверка соединения
	_, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Ошибка подключения к Redis:", err)
		return
	}

	// Кеширование данных
	err = client.Set(ctx, "username", "john_doe", 0).Err()
	if err != nil {
		fmt.Println("Ошибка при кешировании данных:", err)
		return
	}

	err = client.Set(ctx, "age", 25, 0).Err()
	if err != nil {
		fmt.Println("Ошибка при кешировании данных:", err)
		return
	}
	err = client.Set(ctx, "city", "LA", 0).Err()
	if err != nil {
		fmt.Println("Ошибка при кешировании данных:", err)
		return
	}

	// Получение данных из кеша
	username, err := client.Get(ctx, "username").Result()
	if err != nil {
		fmt.Println("Ошибка при получении данных из кеша:", err)
		return
	}

	fmt.Println("Имя пользователя из кеша:", username)
}
