package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"google.golang.org/grpc/metadata"

	"github.com/OtusGolang/webinars_practical_part/27-grpc/elections-with-stats/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var lastID uint64

func GetNextID() uint64 {
	curId := atomic.AddUint64(&lastID, 1)
	return curId
}

func GenerateActionId() string {
	curId := GetNextID()
	return fmt.Sprintf("%v:%v", time.Now().UTC().Format("20060102150405"), curId)
}

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

		md := metadata.New(nil)
		md.Append("request_id", GenerateActionId())
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		if _, err := client.SubmitVote(ctx, req); err != nil {
			log.Fatal(err)
		}

		log.Printf("vote submitted")
	}
}

func getRequest(reader *bufio.Reader) (*pb.Vote, error) {
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

	return &pb.Vote{
		Passport:    parts[0],
		CandidateId: uint32(id),
		Note:        strings.Join(parts[2:], " "),
		Time:        timestamppb.Now(),
	}, nil
}
