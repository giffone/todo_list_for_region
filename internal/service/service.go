package service

import (
	"context"
	"errors"
	"fmt"
	"todolist/internal/domain"
	"todolist/internal/repository"
	"todolist/pkg/hashkey"
)

type service struct {
	db repository.Db
}

func (s *service) CreateTask(ctx context.Context, r *domain.Request) *domain.Response {
	// make unique key
	key := fmt.Sprintf("%s%s", r.Title, r.ActiveAt)
	// make model
	t := domain.TaskDTO{
		Title:    r.Title,
		ActiveAt: r.ValidDate,
		HashKey:  hashkey.MakeHashKey(key), // create unique hash for task
	}
	if err := s.db.CreateTask(ctx, &t); err != nil {
		if errors.Is(err, domain.ErrAlreadyExist) {
			return &domain.StatusAlreadyExist
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return &domain.StatusTimeOut
		}
		res := domain.StatusIntSrvErr
		res.WrapStatus(err.Error())
		return &res
	}
	return &domain.StatusCreated
}
