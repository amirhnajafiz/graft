package main

import (
	"cmd/proto"
	"cmd/server/handler"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	proto.RegisterCustomerServer(s, &handler.Server{})
	log.Printf("Attemp to listen on: %s", port)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
}
