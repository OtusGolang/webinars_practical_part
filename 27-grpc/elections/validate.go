package elections

import (
	"context"
	"errors"
	"github.com/OtusGolang/webinars_practical_part/27-grpc/elections/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Validator func(req interface{}) error

func UnaryServerRequestValidatorInterceptor(validator Validator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := validator(req); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "%s is rejected by validate middleware. Error: %v", info.FullMethod, err)
		}
		return handler(ctx, req)
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

