package main

import (
	"context"
	"github.com/OtusGolang/webinars_practical_part/27-grpc/elections"
	"google.golang.org/grpc/metadata"
	"log"
	"net"

	"github.com/OtusGolang/webinars_practical_part/27-grpc/elections/pb"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type Service struct {
	pb.UnimplementedElectionsServer
}

func (s *Service) SubmitVote(ctx context.Context, req *pb.Vote) (*empty.Empty, error) {
	requestId := ""
	if ctx != nil {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			ids := md.Get("request_id")
			if len(ids) > 0 {
				requestId = ids[0]
			}
		}
	}
	log.Printf("new vote receive (passport=%s, candidate_id=%d, time=%v, request_id: %v)",
		req.Passport, req.CandidateId, ptypes.TimestampString(req.Time), requestId)

	//if req.Passport == "" || req.CandidateId == 0 {
	//	log.Printf("invalid arguments, skip vote")
	//	return nil, status.Error(codes.InvalidArgument, "passport or candidate_id wrong")
	//}

	log.Printf("vote accepted")
	return &empty.Empty{}, nil
}

func main() {
	lsn, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			elections.UnaryServerRequestValidatorInterceptor(elections.ValidateReq),
			),
		)
	pb.RegisterElectionsServer(server, new(Service))

	log.Printf("starting server on %s", lsn.Addr().String())
	if err := server.Serve(lsn); err != nil {
		log.Fatal(err)
	}
}
