package main

import (
	"cmd/customer"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

const (
	address = "localhost:8080"
)

func createCustomer(client customer.CustomerClient, customer *customer.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create Customer: %v\n", err)
	}
	if resp.Success {
		log.Printf("A new Customer has been added with id: %d\n", resp.Id)
	}
}

func getCustomers(client customer.CustomerClient, filter *customer.CustomerFilter) {
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

func main() {
	var conn *grpc.ClientConn
	var err error

	for {
		conn, err = grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			time.Sleep(2 * time.Second)
			fmt.Printf("Did not connect: %v", err)
		} else {
			break
		}
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Terminated with error: %v", err)
		}
	}(conn)

	client := customer.NewCustomerClient(conn)

	temp := &customer.CustomerRequest{
		Id:    101,
		Name:  "Amir hossein",
		Email: "najafi@gmail.com",
		Phone: "0098913428655",
		Addresses: []*customer.CustomerRequest_Address{
			{
				Street:            "Brad st",
				City:              "New York",
				State:             "New York",
				Zip:               "99810",
				IsShippingAddress: false,
			},
			{
				Street:            "Taylor st",
				City:              "New Jersey",
				State:             "New York",
				Zip:               "99810",
				IsShippingAddress: true,
			},
		},
	}

	createCustomer(client, temp)

	temp = &customer.CustomerRequest{
		Id:    82,
		Name:  "Linda",
		Email: "lindi@gmail.com",
		Phone: "+101 44292",
		Addresses: []*customer.CustomerRequest_Address{
			{
				Street:            "Okyway",
				City:              "Kora",
				State:             "Selenoid",
				Zip:               "12301",
				IsShippingAddress: false,
			},
		},
	}

	createCustomer(client, temp)

	filter := &customer.CustomerFilter{
		Keyword: "",
	}

	getCustomers(client, filter)
}
