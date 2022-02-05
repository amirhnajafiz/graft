package terminal

import (
	"bufio"
	"cmd/customer"
	"cmd/pkg/data"
	"cmd/pkg/endpoint"
	"cmd/pkg/reader"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
)

type Terminal struct {
	Conn   *grpc.ClientConn
	Client customer.CustomerClient
}

func (t *Terminal) Run() {
	flag := false
	read := reader.Reader{
		IO: bufio.NewReader(os.Stdin),
	}

	for {
		fmt.Print("> ")

		command, err := read.Get()

		if err != nil {
			log.Fatalf("input error: %v\n", err)
		}

		switch {
		case command == "CC":
			endpoint.CreateCustomer(t.Client, data.FakeClient())
		case command == "GC":
			filter := &customer.CustomerFilter{
				Keyword: "",
			}

			endpoint.GetCustomers(t.Client, filter)
		case command == "EX":
			flag = true
		}

		fmt.Println()

		if flag {
			return
		}
	}
}
