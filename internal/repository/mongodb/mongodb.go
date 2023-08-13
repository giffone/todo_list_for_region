package mongodb

import (
	"context"
	"time"
	"todolist/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	titleField    = "title"
	activeAtField = "activeAt"
	hashField     = "hash"
)

type storage struct {
	tasks *mongo.Collection
}

func NewDb(c *mongo.Client) *storage {
	return &storage{tasks: c.Database(dbName).Collection(collectionT)}
}

func (s *storage) CreateTask(ctx context.Context, t *domain.TaskDTO) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	_, err := s.tasks.InsertOne(ctx, t)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return domain.ErrAlreadyExist
		}
		return err
	}
	return nil
}

func (s *storage) UpdateTask(ctx context.Context, id string, t *domain.TaskDTO) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// prepare
	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			titleField:    t.Title,
			activeAtField: t.ActiveAt,
			hashField:     t.HashKey,
		},
	}
	// update
	result, err := s.tasks.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	// check result
	if result.ModifiedCount == 0 {
		return domain.ErrNotFound
	}
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
