package main

import (
	"errors"

	pb "github.com/OtusGolang/webinars_practical_part/27-grpc/elections-with-admin/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Validator func(req interface{}) error

type recvWrapper struct {
	grpc.ServerStream
	validFunc Validator
	info      *grpc.StreamServerInfo
}

func (s *recvWrapper) RecvMsg(m interface{}) error {
	if err := s.validFunc(m); err != nil {
		return status.Errorf(codes.InvalidArgument, "%s is rejected by validate middleware. Error: %v", s.info.FullMethod, err)
	}
	if err := s.ServerStream.RecvMsg(m); err != nil {
		return err
	}
	return nil
}

func StreamServerRequestValidatorInterceptor(validator Validator) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		wrapper := &recvWrapper{
			ServerStream: ss,
			validFunc:    validator,
			info:         info,
		}
		return handler(srv, wrapper)
	}
}

func ValidateReq(req interface{}) error {
	switch r := req.(type) {
	case *pb.Vote:
		if r.Passport == "" || r.CandidateId == 0 {
			return errors.New("middleware validator: passport or candidate_id wrong")
		}
	}
	return nil
}
