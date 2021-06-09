package main

import (
	"bufio"
	"context"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/OtusGolang/webinars_practical_part/27-grpc/elections/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewElectionsClient(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		req, err := getRequest(reader)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}

		if _, err := client.SubmitVote(context.Background(), req); err != nil {
			log.Fatal(err)
		}

		log.Printf("vote submitted")
	}
}

func getRequest(reader *bufio.Reader) (*pb.SubmitVoteRequest, error) {
	log.Printf("write <passport> <candidate_id> <note>:")
	text, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("wrong input, try again")
	}

	parts := strings.Split(text, " ")
	if len(parts) < 3 {
		return nil, errors.New("wrong input, try again")
	}
	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, errors.New("wrong input, try again")
	}

	return &pb.SubmitVoteRequest{
		Vote: &pb.SubmitVoteRequest_Vote{
			Passport:    parts[0],
			CandidateId: uint32(id),
			Note:        strings.Join(parts[2:], " "),
			Time:        timestamppb.Now(),
		},
	}, nil
}
