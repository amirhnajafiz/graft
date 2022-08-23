package main

import (
	"log"
	"net"

	"github.com/amirhnajafiz/protocol-buffers/proto"
	"google.golang.org/grpc"
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

	proto.RegisterCustomerServer(s, &server{})

	log.Printf("Attemp to listen on: %s", port)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
}
