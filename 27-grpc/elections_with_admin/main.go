package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lsn, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer(
		grpc.ChainStreamInterceptor(
			StreamServerRequestValidatorInterceptor(ValidateReq),
		),
		)
	RegisterElectionsServer(server, NewService())

	log.Printf("starting server on %s", lsn.Addr().String())
	if err := server.Serve(lsn); err != nil {
		log.Fatal(err)
	}
}
