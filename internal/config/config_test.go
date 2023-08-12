package config

import (
	"testing"
	"todolist/internal/domain"
)

func TestServerConf(t *testing.T) {
	confs := []struct{
		addr string
		want error
	}{
		{addr: ":1234", want: nil},
		{addr: "1234", want: nil},
		{addr: "1234:", want: domain.ErrSrvAddrNotValid},
		{addr: ":", want: domain.ErrSrvAddrNotValid},
		{addr: ":1234a", want: domain.ErrSrvAddrNotValid},
		{addr: ":a", want: domain.ErrSrvAddrNotValid},
		{addr: ":1234:", want: domain.ErrSrvAddrNotValid},
		{addr: "", want: domain.ErrSrvAddrEmpty},
	}

	sc := ServerConf{}

	for _, v := range confs {
		sc.Addr = v.addr
		if err := sc.validate(); err != v.want {
			t.Fatalf("want %v but got %v", v.want, err)
		}
	}
}
