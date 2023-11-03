package main

import (
	"context"
	"log"
	"taskmanager/api/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewTaskServiceClient(conn)

	// Create a new task
	createResponse, err := client.CreateTask(context.Background(), &proto.TaskRequest{
		Title:       "New Task",
		Description: "Description for new task",
		Status:      "Pending",
	})
	if err != nil {
		log.Fatalf("Error creating task: %v", err)
	}
	log.Printf("Task created: %v\n", createResponse)
}
