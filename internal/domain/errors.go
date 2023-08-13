package domain

import "errors"

var (
	ErrAlreadyExist = errors.New("Task already exist")
	ErrAddrEmpty    = errors.New("Address is empty")
	ErrAddrNotValid = errors.New("Address not valid")
)
