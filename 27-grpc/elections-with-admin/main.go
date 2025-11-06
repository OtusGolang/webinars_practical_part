package main

import (
	"log"
	"net"

	pb "github.com/OtusGolang/webinars_practical_part/27-grpc/elections-with-admin/pb"
	"google.golang.org/grpc"
)

func main() {
	lsn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer(
		grpc.ChainStreamInterceptor(
			StreamServerRequestValidatorInterceptor(ValidateReq),
		),
	)
	pb.RegisterElectionsServer(server, NewService())

	log.Printf("starting server on %s", lsn.Addr().String())
	if err := server.Serve(lsn); err != nil {
		log.Fatal(err)
	}
}
