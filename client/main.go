package main

import (
	"cmd/customer"
	"cmd/pkg/data"
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

func do() {
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

	createCustomer(client, data.FakeClient())

	time.Sleep(3 * time.Second)

	filter := &customer.CustomerFilter{
		Keyword: "",
	}

	getCustomers(client, filter)
}

func main() {
	for i := 0; i < 5; i++ {
		go do()
	}

	time.Sleep(20 * time.Second)
}
