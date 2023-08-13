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
	// prepare model for db
	t := prepareDTO(r)
	// create
	if err := s.db.CreateTask(ctx, t); err != nil {
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

func (s *service) UpdateTask(ctx context.Context, id string, r *domain.Request) *domain.Response {
	// prepare model for db
	t := prepareDTO(r)
	// update
	if err := s.db.UpdateTask(ctx, id, t); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return &domain.StatusNotFound
		}
		res := domain.StatusIntSrvErr
		res.WrapStatus(err.Error())
		return &res
	}
	return &domain.StatusOK
}

func (s *service) DeleteTask(ctx context.Context, id string) *domain.Response {
	// delete
	if err := s.db.DeleteTask(ctx, id); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return &domain.StatusNotFound
		}
		res := domain.StatusIntSrvErr
		res.WrapStatus(err.Error())
		return &res
	}
	return &domain.StatusOK
}

func prepareDTO(r *domain.Request) *domain.TaskDTO {
	// make unique key
	key := fmt.Sprintf("%s%s", r.Title, r.ActiveAt)
	// make model
	return &domain.TaskDTO{
		Title:    r.Title,
		ActiveAt: r.ValidDate,
		HashKey:  hashkey.MakeHashKey(key), // create unique hash for task
	}
}
