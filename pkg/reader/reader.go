package reader

import (
	"bufio"
	"strings"
)

type Reader struct {
	IO *bufio.Reader
}

var delim = "\n"

func (r *Reader) Get() (string, error) {
	text, err := r.IO.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.Trim(text, delim), nil
}
