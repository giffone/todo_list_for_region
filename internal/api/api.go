package api

import (
	"context"
	"todolist/internal/domain"

	"github.com/labstack/echo/v4"
)

type Service interface {
	CreateTask(ctx context.Context, task *domain.Request) *domain.Response
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
	t := domain.Request{}
	var err error
	// parse data
	if err = c.Bind(&t); err != nil {
		res := domain.StatusInvalidReqBody
		res.WrapStatus(err.Error())
		return c.JSON(res.Code, res)
	}
	// validate data
	if err = t.Validate(); err != nil {
		res := domain.StatusInvalidData
		res.WrapStatus(err.Error())
		return c.JSON(res.Code, res)
	}
	// create task in db
	res := h.svc.CreateTask(c.Request().Context(), &t)
	return c.JSON(res.Code, res)
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
