package main

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"taskmanager/api/proto"
	"taskmanager/internal/task"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	taskService := task.NewService()
	taskHandler := task.NewHandler(taskService)

	proto.RegisterTaskServiceServer(s, taskHandler)
	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Println("Server is listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
