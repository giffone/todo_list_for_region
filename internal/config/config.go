package config

import (
	"fmt"
	"todolist/internal/domain"
)

type ServerConf struct {
	Addr string
}

func (s *ServerConf) validate() error {
	if s.Addr == "" {
		return domain.ErrSrvAddrEmpty
	}
	dg, dt :=0, 0
	for i := 0; i < len(s.Addr); i++ {
		if s.Addr[i] < 0 || s.Addr[i] > '9' {
			if s.Addr[i] == ':' {
				dt = i
				continue
			}
			return domain.ErrSrvAddrNotValid
		}
		dg++
	}
	if dg == 0 || dt != 0 {
		return domain.ErrSrvAddrNotValid
	}
	if s.Addr[0] != ':' {
		s.Addr = fmt.Sprintf(":%s", s.Addr)
	}
	return nil
}
