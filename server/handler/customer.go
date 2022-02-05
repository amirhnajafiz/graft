package handler

import (
	"cmd/proto"
	"context"
	"log"
	"strings"
)

func (s *Server) CreateCustomer(ctx context.Context, in *proto.CustomerRequest) (*proto.CustomerResponse, error) {
	s.savedCustomers = append(s.savedCustomers, in)

	log.Println("Customer created")

	return &proto.CustomerResponse{Id: in.Id, Success: true}, nil
}

func (s *Server) GetCustomers(filter *proto.CustomerFilter, stream proto.Customer_GetCustomersServer) error {
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
