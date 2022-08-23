package main

import (
	"context"
	"log"
	"net"
	"strings"

	"github.com/amirhnajafiz/protocol-buffers/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

// our server type which manages the customers.
// notice that server has two methods that we defined
// in proto file.
type server struct {
	savedCustomers []*proto.CustomerRequest
}

// CreateCustomer
// manages to create a new customer.
func (s *server) CreateCustomer(_ context.Context, in *proto.CustomerRequest) (*proto.CustomerResponse, error) {
	s.savedCustomers = append(s.savedCustomers, in)

	log.Println("Customer created")

	return &proto.CustomerResponse{Id: in.Id, Success: true}, nil
}

// GetCustomers
// manages to list customers and send them over an stream.
func (s *server) GetCustomers(filter *proto.CustomerFilter, stream proto.Customer_GetCustomersServer) error {
	for _, savedCustomer := range s.savedCustomers {
		if filter.Keyword != "" {
			if !strings.Contains(savedCustomer.Name, filter.Keyword) {
				continue
			}
		}

		if err := stream.Send(savedCustomer); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	// first we need to create a listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	// initializing a grpc server
	s := grpc.NewServer()

	// register that server with our server type
	proto.RegisterCustomerServer(s, &server{})

	log.Printf("Attemp to listen on: %s", port)

	// serve the grpc server with created listener
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
}
