package terminal

import (
	"bufio"
	"cmd/pkg/reader"
	"fmt"
	"log"
	"os"
)

func Run() {
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
			// Create customer
		case command == "GC":
			// Get customers
		case command == "EX":
			// Terminate program
			flag = true
		}

		if flag {
			return
		}
	}
}
