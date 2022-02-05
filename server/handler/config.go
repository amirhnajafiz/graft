package handler

import (
	"cmd/proto"
)

type Server struct {
	savedCustomers []*proto.CustomerRequest
}
