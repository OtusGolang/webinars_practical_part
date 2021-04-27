package api

import (
	"context"
	"net"
	"time"

	"github.com/OtusGolang/webinars_practical_part/19-clean-architecture/internal/domain/entities"
	"github.com/OtusGolang/webinars_practical_part/19-clean-architecture/internal/domain/errors"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

type eventUsecaseInterface interface {
	CreateEvent(ctx context.Context, owner, title, text string, startTime *time.Time, endTime *time.Time) (*entities.Event, error)
}

type CalendarServer struct {
	EventUsecases eventUsecaseInterface
}

// implements CalendarServiceServer
func (cs *CalendarServer) CreateEvent(ctx context.Context, req *CreateEventRequest) (*CreateEventResponse, error) {
	owner := ""
	if o := ctx.Value("owner"); o != nil {
		owner, _ = o.(string)
	}

	startTime := (*time.Time)(nil)
	if req.GetStartTime() != nil {
		st, err := ptypes.Timestamp(req.GetStartTime())
		if err != nil {
			return nil, err
		}
		startTime = &st
	}

	endTime := (*time.Time)(nil)
	if req.GetEndTime() != nil {
		et, err := ptypes.Timestamp(req.GetEndTime())
		if err != nil {
			return nil, err
		}
		endTime = &et
	}

	event, err := cs.EventUsecases.CreateEvent(ctx, owner, req.GetTitle(), req.GetText(), startTime, endTime)
	if err != nil {
		if berr, ok := err.(errors.EventError); ok {
			resp := &CreateEventResponse{
				Result: &CreateEventResponse_Error{
					Error: string(berr),
				},
			}
			return resp, nil
		} else {
			return nil, err
		}
	}
	protoEvent := &Event{
		Id:    event.ID.String(),
		Title: event.Title,
		Text:  event.Text,
		Owner: event.Owner,
	}
	if protoEvent.StartTime, err = ptypes.TimestampProto(*event.StartTime); err != nil {
		return nil, err
	}
	if protoEvent.EndTime, err = ptypes.TimestampProto(*event.EndTime); err != nil {
		return nil, err
	}
	resp := &CreateEventResponse{
		Result: &CreateEventResponse_Event{
			Event: protoEvent,
		},
	}
	return resp, nil
}

func (cs *CalendarServer) Serve(addr string) error {
	s := grpc.NewServer()
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	RegisterCalendarServiceServer(s, cs)
	return s.Serve(l)
}
