package endpoint

import (
	"cmd/proto"
	"context"
	"io"
	"log"
)

func CreateCustomer(client proto.CustomerClient, customer *proto.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create Customer: %v\n", err)
	}
	if resp.Success {
		log.Printf("A new Customer has been added with id: %d\n", resp.Id)
	}
}

func GetCustomers(client proto.CustomerClient, filter *proto.CustomerFilter) {
	stream, err := client.GetCustomers(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get customers: %v\n", err)
	}
	for {
		rec, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetCustomers(_) = _, %v\n", client, err)
		}
		log.Printf("Customer: %v\n", rec)
	}
}
