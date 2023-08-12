package api

import "github.com/labstack/echo/v4"


type Handlers struct {
	e *echo.Echo
}

func (h *Handlers) CreateTask(c echo.Context) error {
	return nil
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