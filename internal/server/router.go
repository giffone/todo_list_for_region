package server

import (
	"todolist/internal/api"

	"github.com/labstack/echo/v4"
)

type Server struct {
	srv *echo.Echo
}

func New() *Server {
	srv := Server{
		srv: echo.New(),
	}
	return &srv
}

func (s *Server) registerHandlers(h *api.Handlers)  {
	g := s.srv.Group("/api/todo-list")
	g.POST("/tasks",h.CreateTask)
	g.PUT("/tasks/:id",h.UpdateTask)
	g.DELETE("/tasks/:id",h.DeleteTask)
	g.PUT("/tasks/:id/done",h.DoneTask)
	g.GET("/tasks",h.GetTasks)
}