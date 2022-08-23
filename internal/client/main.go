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
)

func main() {
	var (
		conn *grpc.ClientConn
		err  error
	)

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

	client := proto.NewCustomerClient(conn)
	tr := terminal.Terminal{
		Conn:   conn,
		Client: client,
	}

	tr.Run()
}
