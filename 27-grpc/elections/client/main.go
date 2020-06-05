package main

import (
	"bufio"
	context "context"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

func getRequest(reader *bufio.Reader) (*Vote, error) {
	log.Printf("Write <passport> <candidate_id> <note>:")
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

	return &Vote{
		Passport:    parts[0],
		CandidateId: uint32(id),
		Note:        strings.Join(parts[2:], " "),
		Time:        ptypes.TimestampNow(),
	}, nil
}

func main() {
	addr := "localhost:50051"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := NewElectionsClient(conn)

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
