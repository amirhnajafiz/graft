package main

import (
	"context"
	"log"
	"strings"

	"github.com/amirhnajafiz/protocol-buffers/proto"
)

type server struct {
	savedCustomers []*proto.CustomerRequest
}

func (s *server) CreateCustomer(_ context.Context, in *proto.CustomerRequest) (*proto.CustomerResponse, error) {
	s.savedCustomers = append(s.savedCustomers, in)

	log.Println("Customer created")

	return &proto.CustomerResponse{Id: in.Id, Success: true}, nil
}

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
