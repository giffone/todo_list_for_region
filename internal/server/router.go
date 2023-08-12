package server

import (
	"todolist/internal/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func newRouter(h *api.Handlers) *echo.Echo {
	e:= echo.New()

	// set middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// register handlers
	g := e.Group("/api/todo-list")
	g.POST("/tasks", h.CreateTask)
	g.PUT("/tasks/:id", h.UpdateTask)
	g.DELETE("/tasks/:id", h.DeleteTask)
	g.PUT("/tasks/:id/done", h.DoneTask)
	g.GET("/tasks", h.GetTasks)
	return e
}
