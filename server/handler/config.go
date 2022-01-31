package handler

import "cmd/customer"

type Server struct {
	savedCustomers []*customer.CustomerRequest
}
