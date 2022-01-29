package client

import (
	"cmd/customer"
	"context"
	"io"
	"log"
)

const (
	address = "localhost:8080"
)

func createCustomer(client customer.CustomerClient, customer *customer.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	if resp.Success {
		log.Printf("A new Customer has been added with id: %d", resp.Id)
	}
}

func getCustomers(client customer.CustomerClient, filter *customer.CustomerFilter) {
	stream, err := client.GetCustomers(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get customers: %v", err)
	}
	for {
		rec, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetCustomers(_) = _, %v", client, err)
		}
		log.Printf("Customer: %v", rec)
	}
}
