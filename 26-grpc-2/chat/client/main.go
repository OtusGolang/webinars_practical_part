package main

import (
	"context"
	"log"
	"time"

	chatproto "chat/api/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := chatproto.NewChatServiceClient(conn)

	stream, err := c.Chat(context.Background())
	if err != nil {
		log.Fatalf("error creating stream: %v", err)
	}
	for i := 0; i < 100; i++ {
		// Send a message to the server.
		msg := &chatproto.Message{Content: "Hello, server!"}
		if err := stream.Send(msg); err != nil {
			log.Fatalf("failed to send: %v", err)
		}

		// Receive a message from the server.
		response, err := stream.Recv()
		if err != nil {
			log.Fatalf("error receiving: %v", err)
		}
		log.Printf("Received: %s", response.Content)
		time.Sleep(1 * time.Second)
	}

}
