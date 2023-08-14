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
	env  *env
}

func New(h *api.Handlers) *Server {
	srv := Server{
		quit: make(chan os.Signal, 1),
		e:    newRouter(h),
	}
	// register notifiers
	signal.Notify(srv.quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	return &srv
}

func (s *Server) Run() {
	// stop envs - mongodb etc.
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		s.env.Stop(ctx)
	}()

	// start router
	go func() {
		if err := s.e.Start(s.cfg.Addr); err != nil && err != http.ErrServerClosed {
			log.Printf("server start error: %s\n", err.Error())
			s.quit <- syscall.SIGQUIT
		}
	}()

	// wait notifiers [ctrl-c]
	<-s.quit
	log.Println("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// stop router
	if err := s.e.Shutdown(ctx); err != nil {
		log.Printf("server stopping error: %s\n", err.Error())
		return
	}
	log.Println("server stopped successfully")
}
