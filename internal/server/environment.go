package server

import (
	"context"
	"log"
	"time"
	"todolist/internal/config"
	"todolist/internal/repository"
	"todolist/internal/repository/mongodb"
)

type env struct {
	storage repository.Storage
}

func newEnvorinment(cfg *config.DbConf) (*env, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return &env{storage: mongodb.NewClient(ctx, cfg)}, nil
}

func (e *env) Stop(ctx context.Context) {
	log.Println("stopping envorinments")
	e.storage.Stop(ctx)
	// etc ...
}
