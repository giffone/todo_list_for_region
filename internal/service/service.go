package service

import (
	"todolist/internal/domain"
	"todolist/internal/repository"
	"todolist/pkg/hashkey"
)

type service struct {
	db repository.Db
}

func (s *service) CreateTask(t *domain.Task) error {
	// create unique hash for task
	t.HashKey = hashkey.MakeHashKey(t.Title, t.ActiveAt)
	s.db.CreateTask()
	return nil
}
