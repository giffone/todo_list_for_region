package server

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	"todolist/internal/api"
	"todolist/internal/config"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e   *echo.Echo
	cfg *config.ServerConf
	env *env
}

func New(h *api.Handlers) *Server { return &Server{e: newRouter(h)} }

func (s *Server) Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		// stop envs - mongodb etc
		s.env.Stop(ctx)
	}()

	// start router
	go func() {
		if err := s.e.Start(s.cfg.Addr); err != nil && err != http.ErrServerClosed {
			log.Printf("server start error: %s\n", err.Error())
			cancel()
		}
	}()

	// wait system notifiers or cancel func
	<-ctx.Done()
	log.Println("stopping server")

	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()

	// stop router
	if err := s.e.Shutdown(ctx2); err != nil {
		log.Printf("server stopping error: %s\n", err.Error())
		return
	}
	log.Println("server stopped successfully")
}
