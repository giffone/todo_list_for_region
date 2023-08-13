package mongodb

import (
	"context"
	"todolist/internal/domain"

	"go.mongodb.org/mongo-driver/mongo"
)



type storage struct {
	tasks *mongo.Collection
}

func NewDb(c *mongo.Client) *storage {
	return &storage{tasks: c.Database(dbName).Collection(collectionT)}
}

func (s *storage) CreateTask(ctx context.Context, t *domain.TaskDTO) error {
	_, err := s.tasks.InsertOne(ctx, t)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return domain.ErrAlreadyExist
		}
		return err
	}
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
