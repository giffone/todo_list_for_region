package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todolist/internal/api"
	"todolist/internal/config"

	"github.com/labstack/echo/v4"
)

type Server struct {
	quit chan os.Signal
	e    *echo.Echo
	cfg  *config.ServerConf
	env *env
}

func New(h *api.Handlers) *Server {
	srv := Server{
		quit: make(chan os.Signal, 1),
		e:    newRouter(h),
	}
	// register notifiers
	signal.Notify(srv.quit, syscall.SIGINT, syscall.SIGTERM)
	return &srv
}

func (s *Server) Run() {
	// stop envs - mongodb etc.
	defer s.env.Stop()

	// start router
	go func() {
		if err := s.e.Start(s.cfg.Addr); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server start error: %s", err.Error())
		}
	}()
	
	// wait notifiers
	<-s.quit
	log.Println("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 5 *time.Second)
	defer cancel()

	// stop router
	if err := s.e.Shutdown(ctx); err != nil {
		log.Fatalf("server stopping error: %s", err.Error())
	}
	log.Println("server stopped successfully")
}
