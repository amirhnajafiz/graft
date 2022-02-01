package reader

import "bufio"

type Reader struct {
	IO bufio.Reader
}

var delim byte = 1

func (r *Reader) Get() (string, error) {
	text, err := r.IO.ReadString(delim)
	if err != nil {
		return "", err
	}

	return text, nil
}
