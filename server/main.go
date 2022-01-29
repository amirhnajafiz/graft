package server

import "cmd/customer"

const (
	port = ":8080"
)

type Server struct {
	savedCustomers []*customer.CustomerRequest
}
