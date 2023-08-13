package mongodb

import (
	"context"
	"time"
	"todolist/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// fields
	idField       = "id"
	titleField    = "title"
	activeAtField = "activeAt"
	statusField   = "status"
	hashField     = "hash"
	// sort
	asc  = 1
	desc = -1
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
	filter := bson.M{idField: id}
	update := bson.M{
		"$set": bson.M{
			titleField:    t.Title,
			activeAtField: t.ActiveAt,
			// statusField:   t.Status.String(),
			hashField: t.HashKey,
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

func (s *storage) DeleteTask(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// prepare
	filter := bson.M{idField: id}
	// delete
	result, err := s.tasks.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	// check result
	if result.DeletedCount == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (s *storage) DoneTask(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// prepare
	filter := bson.M{idField: id}
	update := bson.M{"$set": bson.M{statusField: domain.StatusDone.String()}}
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

func (s *storage) GetTasks(ctx context.Context, status string) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// prepare
	var filter bson.M
	if status == domain.StatusDone.String() {
		filter = bson.M{statusField: status}
	} else {
		filter = bson.M{
			statusField:   status,
			activeAtField: bson.M{"$lte": time.Now()},
		}
	}
	// count tasks
	count, err := s.tasks.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, domain.ErrNotFound
	}

	// sort
	opt := options.Find()
	opt.SetSort(bson.M{activeAtField: asc})

	// get task list
	cur, err := s.tasks.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	tasksList := make([]domain.Task, 0, count)
	// iterate and read
	for cur.Next(ctx) {
		var task domain.Task
		if err := cur.Decode(&task); err != nil {
			return nil, err
		}
		tasksList = append(tasksList, task)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	return tasksList, nil
}
