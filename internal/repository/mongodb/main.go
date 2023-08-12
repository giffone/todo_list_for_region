package mongodb

import (
	"context"
	"log"
	"todolist/internal/config"
	"todolist/internal/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const driver = "mongodb"

type cli struct {
	s *storage
}

func NewClient(ctx context.Context, cfg *config.DbConf) repository.Storage {
	if cfg.Driver != driver {
		log.Fatalln("mongodb: addr-connection not valid")
	}
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Addr))
	if err != nil {
		log.Fatalf("mongodb: connection error: %s", err.Error())
	}
	return &cli{s: NewDb(c)}
}

func (c *cli) Stop(ctx context.Context) {
	if c.s.db == nil {
		return
	}
	if err := c.s.db.Disconnect(ctx); err != nil {
		log.Printf("mongodb: stopping error: %s", err.Error())
	} else {
		log.Println("mongodb: stopping successfully")
	}
}

func (c *cli) Methods() repository.Db {
	return c.s
}
