package main

import (
	"context"
	"log"
	"net"

	"github.com/OtusGolang/webinars_practical_part/27-grpc/elections/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	pb.UnimplementedElectionsServer
}

func (s *Service) SubmitVote(ctx context.Context, req *pb.SubmitVoteRequest) (*pb.SubmitVoteResponse, error) {
	vote := req.Vote
	if vote == nil {
		return nil, status.Error(codes.InvalidArgument, "vote is not specified")
	}

	log.Printf("new vote receive (passport=%s, candidate_id=%d, time=%v)",
		vote.Passport, vote.CandidateId, vote.Time.AsTime())

	if vote.Passport == "" || vote.CandidateId == 0 {
		log.Printf("invalid arguments, skip vote")
		return nil, status.Error(codes.InvalidArgument, "passport or candidate_id wrong")
	}

	log.Printf("vote accepted")
	return &pb.SubmitVoteResponse{}, nil
}

func main() {
	lsn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pb.RegisterElectionsServer(server, new(Service))

	log.Printf("starting server on %s", lsn.Addr().String())
	if err := server.Serve(lsn); err != nil {
		log.Fatal(err)
	}
}
