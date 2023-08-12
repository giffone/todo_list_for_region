package server

import (
	"context"
	"log"
	"time"
	"todolist/internal/config"
	"todolist/internal/repository/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
)

type env struct {
	db *mongo.Client
}

func newEnvorinment(cfg *config.DbConf) (*env, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return &env{db: mongodb.New(ctx, cfg)}, nil
}

func (e *env) Stop(ctx context.Context) {
	log.Println("stopping envorinments")

	if e.db != nil {
		if err := e.db.Disconnect(ctx); err != nil {
			log.Printf("stopping db error: %s", err.Error())
		} else {
			log.Println("stopping db successfully")
		}
	}
}
