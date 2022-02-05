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
	"time"
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
		fmt.Print(time.Now().Format("2006-01-02 15:04:05"))
		fmt.Print(" >> ")

		command, err := read.Get()

		if err != nil {
			log.Fatalf("input error: %v\n", err)
		}

		switch {
		case command == "create":
			endpoint.CreateCustomer(t.Client, data.FakeClient())
		case command == "list":
			filter := &customer.CustomerFilter{
				Keyword: "",
			}

			endpoint.GetCustomers(t.Client, filter)
		case command == "exit":
			flag = true
		}

		fmt.Println()

		if flag {
			return
		}
	}
}
