package domain

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrAlreadyExist = errors.New("Task already exist")
	ErrNotFound     = errors.New("Task not found")
	ErrAddrEmpty    = errors.New("Address is empty")
	ErrAddrNotValid = errors.New("Address not valid")
)

type Response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

func (r *Response) WrapStatus(str string) {
	if r.Status == "" {
		r.Status = str
		return
	}
	r.Status = fmt.Sprintf("%s: %s", r.Status, str)
}

var (
	StatusOK = Response{
		Code:   http.StatusOK,
		Status: "OK",
	}
	StatusCreated = Response{
		Code:   http.StatusCreated,
		Status: "Created",
	}
	StatusIntSrvErr = Response{
		Code:   http.StatusInternalServerError,
		Status: "",
	}
	StatusTimeOut = Response{
		Code:   http.StatusRequestTimeout,
		Status: "Try later",
	}
	StatusAlreadyExist = Response{
		Code:   http.StatusNoContent,
		Status: ErrAlreadyExist.Error(),
	}
	StatusNotFound = Response{
		Code:   http.StatusNotFound,
		Status: ErrNotFound.Error(),
	}
	StatusInvalidReqBody = Response{
		Code:   http.StatusBadRequest,
		Status: "Invalid request body",
	}
	StatusInvalidData = Response{
		Code:   http.StatusBadRequest,
		Status: "Invalid data",
	}
)