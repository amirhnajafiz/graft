package terminal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	pgrpc "github.com/amirhnajafiz/protocol-buffers/internal/client/ports/grpc"
	"github.com/amirhnajafiz/protocol-buffers/pkg/faker"
	"github.com/amirhnajafiz/protocol-buffers/pkg/reader"
	"github.com/amirhnajafiz/protocol-buffers/proto"
	"google.golang.org/grpc"
)

type Terminal struct {
	Conn   *grpc.ClientConn
	Client proto.CustomerClient
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
			pgrpc.CreateCustomer(t.Client, faker.FakeClient())
		case command == "list":
			filter := &proto.CustomerFilter{
				Keyword: "",
			}

			pgrpc.GetCustomers(t.Client, filter)
		case command == "exit":
			flag = true
		}

		fmt.Println()

		if flag {
			return
		}
	}
}
