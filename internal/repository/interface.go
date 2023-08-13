package repository

import (
	"context"
	"todolist/internal/domain"
)

type Db interface {
	CreateTask(ctx context.Context, t *domain.TaskDTO) error
	UpdateTask() error
	DeleteTask() error
	DoneTask() error
	GetTasks() error
}

type Storage interface {
	Stop(ctx context.Context)
	Methods() Db
}
