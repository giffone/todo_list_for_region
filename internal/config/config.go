package config

import (
	"fmt"
	"strings"
	"todolist/internal/domain"
)

type ServerConf struct {
	Addr string
}

func (s *ServerConf) validate() error {
	if s.Addr == "" {
		return domain.ErrAddrEmpty
	}
	if err := portValidate(s.Addr); err != nil {
		return err
	}
	if s.Addr[0] != ':' {
		s.Addr = fmt.Sprintf(":%s", s.Addr)
	}
	return nil
}

type DbConf struct {
	Addr string
	Driver string
}

func (d *DbConf) validate() error {
	if d.Addr == "" {
		return domain.ErrAddrEmpty
	}
	a := strings.Split(d.Addr,":")
	if len(a) < 2 {
		return domain.ErrAddrNotValid
	}

	d.Driver = a[0]

	if err := portValidate(a[len(a)-1]); err != nil {
		return err
	}
	return nil
}

func portValidate(port string) error {
	dg, dt := 0, 0
	for i := 0; i < len(port); i++ {
		if port[i] < 0 || port[i] > '9' {
			if port[i] == ':' {
				dt = i
				continue
			}
			return domain.ErrAddrNotValid
		}
		dg++
	}
	if dg == 0 || dt != 0 {
		return domain.ErrAddrNotValid
	}
	return nil
}
