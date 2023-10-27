package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	// Подключение к MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Пинг сервера для проверки соединения
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Подключено к MongoDB!")

	// Создание или переключение на базу данных
	dbName := "mydb"
	db := client.Database(dbName)

	// Создание коллекции
	collectionName := "mycollection"
	collection := db.Collection(collectionName)

	// Вставка документа
	_, err = collection.InsertOne(ctx, bson.M{"name": "John Doe"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Добавлен документ в коллекцию.")

	// Поиск документа
	var result bson.M
	err = collection.FindOne(ctx, bson.M{"name": "John Doe"}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Найден документ:", result)

	// Создание индекса на поле
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: 1}}, // Создание восходящего индекса на поле "name"
		Options: options.Index().SetUnique(true), // Сделать индекс уникальным
	}

	// Создание индекса в коллекции
	_, err = collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Создан уникальный индекс на поле 'name'.")

	// Удаление коллекции
	if err := collection.Drop(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Коллекция удалена.")

	// Удаление базы данных
	if err := client.Database(dbName).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("База данных удалена.")

	// Закрытие соединения с MongoDB
	if err := client.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Отключено от MongoDB.")
}
