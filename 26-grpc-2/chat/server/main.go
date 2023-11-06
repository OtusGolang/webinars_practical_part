package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	chatproto "chat/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	chatproto.UnimplementedChatServiceServer
}

func (s *server) Chat(stream chatproto.ChatService_ChatServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// Echo the received message back to the client.
		if err := stream.Send(msg); err != nil {
			return err
		}
		fmt.Println(msg)
	}
}

func (s *server) StartCount(_ *emptypb.Empty, stream chatproto.ChatService_StartCountServer) error {
	cnt := 1
	for cnt < 1000 {
		// Echo the received message back to the client.
		cnt++
		if err := stream.Send(&chatproto.Message{Content: strconv.Itoa(cnt)}); err != nil {
			return err
		}
		fmt.Println(cnt)
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	chatproto.RegisterChatServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
