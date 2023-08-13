package api

import (
	"fmt"
	"net/http"
	"todolist/internal/domain"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	e *echo.Echo
	//svc
}

func New(e *echo.Echo) *Handlers {
	return &Handlers{e: e}
}

func (h *Handlers) CreateTask(c echo.Context) error {
	t := domain.List{}
	var err error

	if err = c.Bind(&t); err != nil {
		e := fmt.Sprintf("Invalid request body: %s", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": e})
	}
	if err = t.Validate(); err != nil {
		e := fmt.Sprintf("Invalid data: %s", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"error": e})
	}
	response := map[string]interface{}{}
	// if already exist - 204
	return c.JSON(http.StatusCreated, response)
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
