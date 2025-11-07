package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/OtusGolang/webinars_practical_part/27-grpc/elections-with-admin/pb"
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

	log.Printf("starting vote stream...")
	voteStream, err := client.Internal(context.Background())
	if err != nil {
		log.Fatal("error on get internal stream", err)
	}

	// go routine to receive admin messages
	go func() {
		for {
			statsVote, err := voteStream.Recv()
			if err != nil {
				if err == io.EOF {
					log.Printf("vote stream ended by server")
					os.Exit(0)
					return
				}
				log.Fatal("error on receive from internal stream", err)
			}

			log.Printf("admin message received")
			v := statsVote.GetVote()
			if v != nil {
				log.Printf("admin vote received: passport=%s candidate_id=%d note=%s time=%s",
					v.Passport, v.CandidateId, v.Note, v.Time.AsTime().Format(time.RFC3339))
			}

			s := statsVote.GetStats()
			if s != nil {
				log.Printf("admin stats received: %v", s.Records)
			}
		}
	}()

	// main routine to send votes
	reader := bufio.NewReader(os.Stdin)
	for {
		vote, err := getRequest(reader)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}

		err = voteStream.Send(vote)
		if err != nil {
			log.Fatal("error on send vote to internal stream", err)
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
