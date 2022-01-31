package handler

import (
	"cmd/customer"
	"context"
	"log"
	"strings"
)

func (s *Server) CreateCustomer(ctx context.Context, in *customer.CustomerRequest) (*customer.CustomerResponse, error) {
	s.savedCustomers = append(s.savedCustomers, in)

	log.Println("Customer created")

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
