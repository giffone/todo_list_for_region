package api

import (
	"context"
	"todolist/internal/domain"

	"github.com/labstack/echo/v4"
)

type Service interface {
	CreateTask(ctx context.Context, task *domain.Request) *domain.Response
	UpdateTask(ctx context.Context, id string, r *domain.Request) *domain.Response
	DeleteTask(ctx context.Context, id string) *domain.Response
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
	// parse data from req
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
	t := domain.Request{}
	var err error
	// parse data from req
	id := c.Param("id")
	if err = c.Bind(&t); err != nil {
		res := domain.StatusInvalidReqBody
		res.WrapStatus(err.Error())
		return c.JSON(res.Code, res)
	}
	// validate data
	if id == "" {
		res := domain.StatusInvalidData
		res.WrapStatus("Not valid id")
		return c.JSON(res.Code, res)
	}
	if err = t.Validate(); err != nil {
		res := domain.StatusInvalidData
		res.WrapStatus(err.Error())
		return c.JSON(res.Code, res)
	}
	// update task in db
	res := h.svc.UpdateTask(c.Request().Context(), id, &t)
	return c.JSON(res.Code, res)
}

func (h *Handlers) DeleteTask(c echo.Context) error {
	// parse data from req
	id := c.Param("id")
	// validate data
	if id == "" {
		res := domain.StatusInvalidData
		res.WrapStatus("Not valid id")
		return c.JSON(res.Code, res)
	}
	// delete task in db
	res := h.svc.DeleteTask(c.Request().Context(), id)
	return c.JSON(res.Code, res)
}

func (h *Handlers) DoneTask(c echo.Context) error {
	return nil
}

func (h *Handlers) GetTasks(c echo.Context) error {
	return nil
}
