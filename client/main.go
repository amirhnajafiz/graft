package main

import (
	"cmd/customer"
	terminal2 "cmd/terminal"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:8080"
)

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
	terminal := terminal2.Terminal{
		Conn:   conn,
		Client: client,
	}

	terminal.Run()
}

func main() {
	do()
}
