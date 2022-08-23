package main

import (
	"fmt"
	"log"
	"time"

	"github.com/amirhnajafiz/protocol-buffers/internal/client/terminal"
	"github.com/amirhnajafiz/protocol-buffers/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
	timeout = 2
)

func main() {
	var (
		conn *grpc.ClientConn
		err  error
	)

	// connecting to grpc server
	for {
		conn, err = grpc.Dial(address)
		if err != nil {
			time.Sleep(timeout * time.Second)
			fmt.Printf("Did not connect: %v", err)
		} else {
			break
		}
	}

	// closing our grpc connection
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Terminated with error: %v", err)
		}
	}(conn)

	// now we create a new customer and begin
	client := proto.NewCustomerClient(conn)
	tr := terminal.Terminal{
		Client: client,
	}

	// start terminal
	tr.Run()
}
