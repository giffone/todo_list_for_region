package mongodb

import (
	"context"
	"log"
	"todolist/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const driver = "mongodb"

func New(ctx context.Context, cfg *config.DbConf) *mongo.Client {
	if cfg.Driver != driver {
		log.Fatalln("mongodb: addr-connection not valid")
	}
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Addr))
	if err != nil {
		log.Fatalf("mongodb: connection error: %s", err.Error())
	}
	return db
}
