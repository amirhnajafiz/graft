package data

import (
	"cmd/customer"
	"github.com/bxcodec/faker/v3"
	"math/rand"
)

func FakeClient() *customer.CustomerRequest {
	return &customer.CustomerRequest{
		Id:    int32(rand.Int()),
		Name:  faker.Name(),
		Email: faker.Email(),
		Phone: faker.Phonenumber(),
		Addresses: []*customer.CustomerRequest_Address{
			{
				Street:            faker.FirstName(),
				City:              faker.Word(),
				State:             faker.Word(),
				Zip:               faker.UUIDDigit(),
				IsShippingAddress: rand.Int()%2 == 0,
			},
		},
	}
}
