package main

import (
	"context"
	"log"
	"sync"
	"time"

	pb "github.com/OtusGolang/webinars_practical_part/27-grpc/elections-with-stats/pb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const defaultInterval = 5 * time.Second

type Service struct {
	pb.UnimplementedElectionsServer

	lock     sync.RWMutex
	stats    map[uint32]uint32
	interval time.Duration
}

func NewService() *Service {
	return &Service{
		stats:    make(map[uint32]uint32),
		interval: defaultInterval,
	}
}

func (s *Service) SubmitVote(ctx context.Context, req *pb.Vote) (*empty.Empty, error) {
	log.Printf("new vote receive (passport=%s, candidate_id=%d, time=%v)",
		req.Passport, req.CandidateId, req.Time.AsTime())

	if req.Passport == "" || req.CandidateId == 0 {
		log.Printf("invalid arguments, skip vote")
		return nil, status.Error(codes.InvalidArgument, "passport or candidate_id wrong")
	}

	s.lock.Lock()
	s.stats[req.CandidateId]++
	s.lock.Unlock()

	log.Printf("vote accepted")
	return &empty.Empty{}, nil
}

func (s *Service) GetStats(req *empty.Empty, srv pb.Elections_GetStatsServer) error {
	log.Printf("new stats listener")

L:
	for {
		select {
		case <-srv.Context().Done():
			log.Printf("stats listener disconnected")
			break L

		case <-time.After(s.interval):
			s.lock.RLock()
			stats := make(map[uint32]uint32, len(s.stats))
			for k, v := range s.stats {
				stats[k] = v
			}
			s.lock.RUnlock()

			msg := &pb.Stats{
				Records: stats,
				Time:    timestamppb.Now(),
			}
			if err := srv.Send(msg); err != nil {
				log.Printf("unable to send message to stats listener: %v", err)
				break L
			}

		}
	}

	return nil
}
