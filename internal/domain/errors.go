package domain

import "errors"

var (
	ErrAddrEmpty    = errors.New("address is empty")
	ErrAddrNotValid = errors.New("address not valid")
)
