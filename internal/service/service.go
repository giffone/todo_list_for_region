package service

import (
	"fmt"
	"todolist/internal/domain"
	"todolist/internal/repository"
	"todolist/pkg/hashkey"
)

type service struct {
	db repository.Db
}

func (s *service) CreateTask(r *domain.Request) error {
	// make unique key
	key:= fmt.Sprintf("%s%s", r.Title,r.ActiveAt)
	// make model
	t := domain.TaskDTO{
		Title: r.Title,
		ActiveAt: r.ValidDate,
		HashKey: hashkey.MakeHashKey(key), // create unique hash for task
	}
	return s.db.CreateTask(t)
}
