package main

import (
	"cmd/customer"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

const (
	port = ":8080"
)

type Server struct {
	savedCustomers []*customer.CustomerRequest
}

func (s *Server) CreateCustomer(ctx context.Context, in *customer.CustomerRequest) (*customer.CustomerResponse, error) {
	s.savedCustomers = append(s.savedCustomers, in)
	return &customer.CustomerResponse{Id: in.Id, Success: true}, nil
}

func (s *Server) GetCustomers(filter *customer.CustomerFilter, stream customer.Customer_GetCustomersServer) error {
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
