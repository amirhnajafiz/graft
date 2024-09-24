package storage

import "github.com/amirhnajafiz/bpb/pkg/rpc/logs"

type storage struct {
	logs []*logs.Log
}

func (s *storage) Append(log *logs.Log) {
	s.logs = append(s.logs, log)
}

func (s *storage) Fetch() []*logs.Log {
	return s.logs
}
