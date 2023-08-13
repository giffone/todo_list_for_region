package api

import (
	"errors"
	"fmt"
	"net/http"
	"todolist/internal/domain"
	"todolist/pkg/hashkey"

	"github.com/labstack/echo/v4"
)

type Service interface {
	CreateTask(task *domain.Task) error
	UpdateTask() error
	DeleteTask() error
	DoneTask() error
	GetTasks() error
}

type Handlers struct {
	e   *echo.Echo
	svc Service
}

func New(e *echo.Echo, svc Service) *Handlers {
	return &Handlers{
		e:   e,
		svc: svc,
	}
}

func (h *Handlers) CreateTask(c echo.Context) error {
	t := domain.Task{}
	var err error
	// parse data
	if err = c.Bind(&t); err != nil {
		e := fmt.Sprintf("Invalid request body: %s", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"status": e})
	}
	// validate data
	if err = t.Validate(); err != nil {
		e := fmt.Sprintf("Invalid data: %s", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"status": e})
	}
	//create unique hash
	t.HashKey = hashkey.MakeHashKey(t.Title, t.ActiveAt)
	// create task in db
	err = h.svc.CreateTask(&t)
	if err != nil {
		if errors.Is(err, domain.ErrAlreadyExist) {
			return c.JSON(http.StatusNoContent, map[string]string{"status": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": err.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]string{"status": "ok"})
}

func (h *Handlers) UpdateTask(c echo.Context) error {
	return nil
}

func (h *Handlers) DeleteTask(c echo.Context) error {
	return nil
}

func (h *Handlers) DoneTask(c echo.Context) error {
	return nil
}

func (h *Handlers) GetTasks(c echo.Context) error {
	return nil
}
