package main

import (
	"cmd/customer"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":8080"
)

type Server struct {
	savedCustomers []*customer.CustomerRequest
}

func (s *Server) CreateCustomer(ctx context.Context, in *customer.CustomerRequest) (*customer.CustomerResponse, error) {
	return nil, nil
}

func (s *Server) GetCustomers(filter *customer.CustomerFilter, stream customer.Customer_GetCustomersServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	customer.RegisterCustomerServer(s, &Server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
}