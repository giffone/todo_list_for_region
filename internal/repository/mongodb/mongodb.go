package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type storage struct {
	db *mongo.Client
}

func NewDb(c *mongo.Client) *storage {
	return &storage{db: c}
}

func (s *storage) CreateTask() error {
	return nil
}

func (s *storage) UpdateTask() error {
	return nil
}

func (s *storage) DeleteTask() error {
	return nil
}

func (s *storage) DoneTask() error {
	return nil
}

func (s *storage) GetTasks() error {
	return nil
}
