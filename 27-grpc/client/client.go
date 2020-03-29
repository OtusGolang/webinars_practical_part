package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"otus-examples/otusrpc/homeworkpb"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

func main() {

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := homeworkpb.NewHomeworkCheckerClient(cc)
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	grade, err := c.CheckHomework(ctx, &homeworkpb.CheckHomeworkRequest{Hw: 10, Code: ""})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Deadline exceeded!")
			} else {
				fmt.Printf("undexpected error %s\n", statusErr.Message())
			}
		} else {
			fmt.Printf("Error while calling RPC CheckHomework: %v", err)
		}
	} else {
		println(grade.Grade)
	}

	stream, err := c.CheckAllHomeworks(context.Background(), &homeworkpb.CheckAllHomeworksRequest{Hw: []int32{1, 2, 3, 4, 5}})
	if err != nil {
		log.Fatalf("CheckAllHomeworks err %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error reading stream: %v", err)
		}
		print(msg.Grade)
	}

	requests := []*homeworkpb.SubmitAllHomeworksRequest{
		&homeworkpb.SubmitAllHomeworksRequest{Hw: 1, Code: "first"},
		&homeworkpb.SubmitAllHomeworksRequest{Hw: 2, Code: "second"},
	}
	cstream, err := c.SubmitAllHomeworks(context.Background())
	if err != nil {
		log.Fatalf("err streaming: %v", err)
	}
	for _, req := range requests {
		cstream.Send(req)
	}

	res, err := cstream.CloseAndRecv()
	if err != nil {
		log.Fatalf("err getting resp: %v", err)
	}
	println(res.GetAccepted())

	// bi-directional

	bstream, err := c.RealtimeFeedback(context.Background())
	if err != nil {
		log.Fatalf("%v", err)
	}

	brequests := []*homeworkpb.CheckHomeworkRequest{
		&homeworkpb.CheckHomeworkRequest{
			Hw:   12,
			Code: "some code",
		},
		&homeworkpb.CheckHomeworkRequest{
			Hw:   13,
			Code: "other code",
		},
	}

	waitc := make(chan struct{})
	// we send a bunch of messages to the client (go routine)
	go func() {
		// function to send a bunch of messages
		for _, req := range brequests {
			fmt.Printf("Sending message: %v\n", req)
			bstream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	// we receive a bunch of messages from the client (go routine)
	go func() {
		// function to receive a bunch of messages
		for {
			res, err := bstream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				break
			}
			fmt.Printf("Received: %v\n", res.GetGrade())
		}
		close(waitc)
	}()

	// block until everything is done
	<-waitc
}
