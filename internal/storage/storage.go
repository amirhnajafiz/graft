package storage

import "github.com/amirhnajafiz/bpb/pkg/rpc/logs"

type Storage interface {
	Append(log *logs.Log)
	Fetch() []*logs.Log
}

func NewStorage() Storage {
	s := storage{
		logs: make([]*logs.Log, 0),
	}

	return &s
}
