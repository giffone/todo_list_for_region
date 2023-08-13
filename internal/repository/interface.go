package repository

import (
	"context"
	"todolist/internal/domain"
)

type Db interface {
	CreateTask(ctx context.Context, t *domain.TaskDTO) error
	UpdateTask(ctx context.Context, id string, t *domain.TaskDTO) error
	DeleteTask(ctx context.Context, id string) error
	DoneTask(ctx context.Context, id string) error
	GetTasks(ctx context.Context, status string) ([]domain.Task, error)
}

type Storage interface {
	Stop(ctx context.Context)
	Methods() Db
}
