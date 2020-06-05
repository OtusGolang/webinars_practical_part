package main

import (
	context "context"
	"errors"
	"log"
	"net"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const defaultInterval = 5 * time.Second

type Service struct {
	lock  sync.RWMutex
	stats map[uint32]uint32
}

func (s *Service) getStats() *Stats {
	s.lock.RLock()
	defer s.lock.RUnlock()

	stats := make(map[uint32]uint32)
	for k, v := range s.stats {
		stats[k] = v
	}

	return &Stats{
		Records: stats,
		Time:    ptypes.TimestampNow(),
	}
}

func (s *Service) submitVote(req *Vote) error {
	log.Printf("new vote receive (passport=%s, candidate_id=%d, time=%v)",
		req.Passport, req.CandidateId, ptypes.TimestampString(req.Time))

	if req.Passport == "" || req.CandidateId == 0 {
		return errors.New("invalid arguments, skip vote")
	}

	s.lock.Lock()
	s.stats[req.CandidateId]++
	s.lock.Unlock()

	log.Printf("vote accepted")
	return nil
}

func (s *Service) SubmitVote(ctx context.Context, req *Vote) (*empty.Empty, error) {
	if err := s.submitVote(req); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *Service) Internal(srv Elections_InternalServer) error {
	log.Printf("new internal listener")

	inChan := make(chan *Vote)
	go func() {
		defer close(inChan)

		for {
			req, err := srv.Recv()
			if err != nil {
				log.Printf("unable to read message from internal listener: %v", err)
				return
			}
			inChan <- req
		}
	}()

	stop := false
	for !stop {
		select {
		case req, ok := <-inChan:
			if !ok {
				log.Printf("read loop for internal listener stopped, disconnect it")
				stop = true
				break
			}
			if _, err := s.SubmitVote(srv.Context(), req); err != nil {
				log.Printf("unable to submit vote, skip it, error: %v", err)
				continue
			}
			msg := &StatsVote{
				Body: &StatsVote_Vote{
					Vote: req,
				},
			}
			if err := srv.Send(msg); err != nil {
				log.Printf("unable to send vote to internal listener, disconnect it, error: %v", err)
				stop = true
			}

		case <-time.After(defaultInterval):
			msg := &StatsVote{
				Body: &StatsVote_Stats{
					Stats: s.getStats(),
				},
			}
			if err := srv.Send(msg); err != nil {
				log.Printf("unable to send stats to internal listener, disconnect it, error: %v", err)
				stop = true
			}

		case <-srv.Context().Done():
			log.Printf("stats listener disconnected")
			stop = true
		}
	}

	return nil
}

func main() {
	lsn, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	service := &Service{stats: make(map[uint32]uint32)}
	RegisterElectionsServer(server, service)

	log.Printf("Starting server on %s", lsn.Addr().String())
	if err := server.Serve(lsn); err != nil {
		log.Fatal(err)
	}
}
