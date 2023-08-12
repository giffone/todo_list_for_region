package repository

import "context"

type Db interface {
	CreateTask() error
	UpdateTask() error
	DeleteTask() error
	DoneTask() error
	GetTasks() error
}

type Storage interface {
	Stop(ctx context.Context)
	Methods() Db
}
