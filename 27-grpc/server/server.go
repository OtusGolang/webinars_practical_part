package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc/reflection"

	"otus-examples/otusrpc/homeworkpb"

	"google.golang.org/grpc"
)

type otusServer struct {
}

func (s *otusServer) CheckHomework(ctx context.Context, req *homeworkpb.CheckHomeworkRequest) (*homeworkpb.CheckHomeworkResponse, error) {

	time.Sleep(time.Millisecond * 200)
	if req.GetCode() == "" {
		return nil, errors.New("Empty submission!")
	}
	return &homeworkpb.CheckHomeworkResponse{Grade: 10}, nil
}

func (s *otusServer) CheckAllHomeworks(req *homeworkpb.CheckAllHomeworksRequest, stream homeworkpb.HomeworkChecker_CheckAllHomeworksServer) error {
	for _, hw := range req.Hw {
		res := &homeworkpb.CheckHomeworkResponse{Hw: hw, Grade: 67}
		stream.Send(res)
		time.Sleep(time.Second)
	}

	return nil
}

func (s *otusServer) SubmitAllHomeworks(stream homeworkpb.HomeworkChecker_SubmitAllHomeworksServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&homeworkpb.SubmitAllHomeworksResponse{Accepted: true})
		}
		if err != nil {
			log.Fatalf("error reading client stream: %v", err)
		}
		_ = req
	}
}

func (s *otusServer) RealtimeFeedback(stream homeworkpb.HomeworkChecker_RealtimeFeedbackServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		_ = req
		sendErr := stream.Send(&homeworkpb.CheckHomeworkResponse{Hw: 1, Grade: 5})
		if sendErr != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	homeworkpb.RegisterHomeworkCheckerServer(grpcServer, &otusServer{})
	grpcServer.Serve(lis)
}
