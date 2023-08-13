package domain

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrAlreadyExist = errors.New("Task already exist")
	ErrAddrEmpty    = errors.New("Address is empty")
	ErrAddrNotValid = errors.New("Address not valid")
)

type Response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

func (r *Response) WrapStatus(str string)  {
	if r.Status == "" {
		r.Status = str
		return
	}
	r.Status = fmt.Sprintf("%s: %s", r.Status, str)
}

var (
	StatusCreated = Response{
		Code:   http.StatusCreated,
		Status: "OK",
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
	StatusInvalidReqBody = Response{
		Code:   http.StatusBadRequest,
		Status: "Invalid request body",
	}
	StatusInvalidData = Response{
		Code:   http.StatusBadRequest,
		Status: "Invalid data",
	}
)
