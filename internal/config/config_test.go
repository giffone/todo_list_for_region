package config

import (
	"testing"
	"todolist/internal/domain"
)

func TestServerConf(t *testing.T) {
	confs := []struct {
		addr string
		want error
	}{
		{addr: ":1234", want: nil},
		{addr: "1234", want: nil},
		{addr: "1234:", want: domain.ErrAddrNotValid},
		{addr: ":", want: domain.ErrAddrNotValid},
		{addr: ":1234a", want: domain.ErrAddrNotValid},
		{addr: ":a", want: domain.ErrAddrNotValid},
		{addr: ":1234:", want: domain.ErrAddrNotValid},
		{addr: "", want: domain.ErrAddrEmpty},
	}

	sc := ServerConf{}

	for _, v := range confs {
		sc.Addr = v.addr
		if err := sc.validate(); err != v.want {
			t.Fatalf("want %v but got %v", v.want, err)
		}
	}
}

func TestDbConf(t *testing.T) {
	confs := []struct {
		addr string
		want error
	}{
		{addr: "mongodb://localhost:1234", want: nil},
		{addr: "mongodb://0.0.0.0:1234", want: nil},
		{addr: "mongodb://localhost:", want: domain.ErrAddrNotValid},
		{addr: ":", want: domain.ErrAddrNotValid},
		{addr: "mongodb:", want: domain.ErrAddrNotValid},
		{addr: ":mongodb", want: domain.ErrAddrNotValid},
		{addr: "1234", want: domain.ErrAddrNotValid},
		{addr: ":1234a", want: domain.ErrAddrNotValid},
		{addr: ":1234:", want: domain.ErrAddrNotValid},
		{addr: "", want: domain.ErrAddrEmpty},
	}

	d := DbConf{}

	for _, v := range confs {
		d.Addr = v.addr
		if err := d.validate(); err != v.want {
			t.Fatalf("want %v but got %v", v.want, err)
		}
	}
}
