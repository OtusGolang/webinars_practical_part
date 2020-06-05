package main

import (
	context "context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes"
	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct{}

func (s *Service) SubmitVote(ctx context.Context, req *Vote) (*empty.Empty, error) {
	log.Printf("new vote receive (passport=%s, candidate_id=%d, time=%v)",
		req.Passport, req.CandidateId, ptypes.TimestampString(req.Time))

	if req.Passport == "" || req.CandidateId == 0 {
		log.Printf("invalid arguments, skip vote")
		return nil, status.Error(codes.InvalidArgument, "passport or candidate_id wrong")
	}

	log.Printf("vote accepted")
	return &empty.Empty{}, nil
}

func main() {
	lsn, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	service := &Service{}
	RegisterElectionsServer(server, service)

	log.Printf("Starting server on %s", lsn.Addr().String())
	if err := server.Serve(lsn); err != nil {
		log.Fatal(err)
	}
}
