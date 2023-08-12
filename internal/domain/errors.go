package domain

import "errors"

var (
	ErrSrvAddrEmpty    = errors.New("server address is empty")
	ErrSrvAddrNotValid = errors.New("server address not valid")
)
