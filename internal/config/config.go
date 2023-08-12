package config

import (
	"errors"
	"fmt"
)

type ServerConf struct {
	Addr string
}

func (s *ServerConf) validate() error {
	if s.Addr == "" {
		return errors.New("server address is empty")
	}
	for i := 0; i < len(s.Addr); i++ {
		if s.Addr[i] < 0 && s.Addr[i] > '9' || s.Addr[i] != ':' {
			return errors.New("server address not valid")
		}
	}
	if s.Addr[0] != ':' {
		s.Addr = fmt.Sprintf(":%s", s.Addr)
	}
	return nil
}
