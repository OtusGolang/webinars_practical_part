package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const defaultInterval = 2 * time.Second

var (
	upgrader = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
)

type Response struct {
	Data  interface{} `json:"data"`
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

type VoteRequest struct {
	Passport    string    `json:"passport,omitempty"`
	CandidateId uint32    `json:"candidate_id,omitempty"`
	Note        string    `json:"note,omitempty"`
	Time        time.Time `json:"time,omitempty"`
}

type StatResponse struct {
	Records map[uint32]uint32 `json:"records,omitempty"`
	Time    time.Time         `json:"time,omitempty"`
}

type StatCandidateResponse struct {
	CandidateId uint32    `json:"candidate_id,omitempty"`
	Stat        uint32    `json:"stat"`
	Time        time.Time `json:"time,omitempty"`
}

type Service struct {
	sync.RWMutex
	Stats    map[uint32]uint32
	Interval time.Duration
}

func NewService() *Service {
	return &Service{
		Stats:    make(map[uint32]uint32),
		Interval: defaultInterval,
	}
}

func (s *Service) SubmitVote(w http.ResponseWriter, r *http.Request) {
	resp := &Response{}
	if r.Method != http.MethodPost {
		resp.Error.Message = fmt.Sprintf("method %s not not supported on uri %s", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusMethodNotAllowed)
		WriteResponse(w, resp)
		return
	}
	buf := make([]byte, r.ContentLength)
	_, err := r.Body.Read(buf)
	if err != nil && err != io.EOF {
		resp.Error.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		WriteResponse(w, resp)
		return
	}

	req := &VoteRequest{}
	err = json.Unmarshal(buf, req)
	if err != nil {
		resp.Error.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		WriteResponse(w, resp)
		return
	}

	// validate field
	if req.Passport == "" || req.CandidateId == 0 {
		slog.Warn("invalid arguments, skip vote")
		resp.Error.Message = "passport or candidate_id wrong"
		w.WriteHeader(http.StatusBadRequest)
		WriteResponse(w, resp)
		return
	}

	slog.Info("new vote receive", "passport", req.Passport, "candidate_id", req.CandidateId, "time", req.Time)

	s.Lock()
	s.Stats[req.CandidateId]++
	s.Unlock()

	slog.Info("vote accepted")
	w.WriteHeader(http.StatusOK)
}

func (s *Service) GetStats(w http.ResponseWriter, r *http.Request) {
	resp := &Response{}
	if r.Method != http.MethodGet {
		resp.Error.Message = fmt.Sprintf("method %s not not supported on uri %s", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusMethodNotAllowed)
		WriteResponse(w, resp)
		return
	}
	args := r.URL.Query()
	id := args.Get("candidate_id")
	if len(id) > 0 {
		candidateId, err := strconv.Atoi(id)
		if err != nil {
			resp.Error.Message = fmt.Sprintf("cant parse candidate_id, expect int, got: %s ", id)
			w.WriteHeader(http.StatusBadRequest)
			WriteResponse(w, resp)
			return
		}

		s.Lock()
		stat, ok := s.Stats[uint32(candidateId)]
		s.Unlock()
		slog.Info("candidate found", "found", ok)
		if !ok {
			resp.Error.Message = fmt.Sprintf("candidate with id %d doasn't found", candidateId)
			w.WriteHeader(http.StatusBadRequest)
			WriteResponse(w, resp)
			return
		}

		resp.Data = &StatCandidateResponse{
			CandidateId: uint32(candidateId),
			Stat:        stat,
			Time:        time.Now(),
		}

		w.WriteHeader(http.StatusOK)
		WriteResponse(w, resp)
		return
	}

	s.Lock()
	stats := s.Stats
	s.Unlock()

	resp.Data = &StatResponse{
		Records: stats,
		Time:    time.Now(),
	}

	w.WriteHeader(http.StatusOK)
	WriteResponse(w, resp)
}

// websocket handler
func (s *Service) StatStream(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	go s.writeLoop(conn)
}

func (s *Service) writeLoop(c *websocket.Conn) {
	defer func() {
		slog.Info("web socket closed")
	}()
	for {
		s.Lock()
		r := s.Stats
		s.Unlock()
		stat := &StatResponse{
			Records: r,
			Time:    time.Now(),
		}

		msg, err := json.Marshal(stat)
		if err != nil {
			c.Close()
			break
		}

		err = c.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			c.Close()
			break
		}
		time.Sleep(s.Interval)
	}
}

func WriteResponse(w http.ResponseWriter, resp *Response) {
	resBuf, err := json.Marshal(resp)
	if err != nil {
		slog.Error("responce marshal error", "err", err)
	}
	_, err = w.Write(resBuf)
	if err != nil {
		slog.Error("responce marshal error", "err", err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}
