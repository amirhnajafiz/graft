package reader

import (
	"bufio"
)

type Reader struct {
	IO *bufio.Reader
}

func (r *Reader) Get() (string, error) {
	text, err := r.IO.ReadString('\n')
	if err != nil {
		return "", err
	}

	return text, nil
}
