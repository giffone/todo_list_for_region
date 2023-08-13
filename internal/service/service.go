package service

import (
	"context"
	"errors"
	"fmt"
	"time"
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

func (s *service) DoneTask(ctx context.Context, id string) *domain.Response {
	// update
	if err := s.db.DoneTask(ctx, id); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return &domain.StatusNotFound
		}
		res := domain.StatusIntSrvErr
		res.WrapStatus(err.Error())
		return &res
	}
	return &domain.StatusOK
}

func (s *service) GetTasks(ctx context.Context, status string) *domain.ResponseList {
	// get list
	tasks, err := s.db.GetTasks(ctx, status); 
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return &domain.ResponseList{Response: domain.StatusNotFound} 
		}
		res := domain.ResponseList{Response: domain.StatusIntSrvErr}
		res.Response.WrapStatus(err.Error())
		return &res
	}
	// add weekend
	for i:= 0; i<len(tasks);i++ {
		w := tasks[i].ActiveAt.Weekday()
		if w == time.Saturday || w == time.Sunday {
			tasks[i].Title = fmt.Sprintf("ВЫХОДНОЙ - %s", tasks[i].Title)
		}
	}
	return &domain.ResponseList{Response:domain.StatusOK, List: tasks}
}

func prepareDTO(r *domain.Request) *domain.TaskDTO {
	// make unique key
	key := fmt.Sprintf("%s%s", r.Title, r.ActiveAt)
	// make model
	return &domain.TaskDTO{
		Title:    r.Title,
		ActiveAt: r.ValidDate,
		Status:   domain.StatusActive,
		HashKey:  hashkey.MakeHashKey(key), // create unique hash for task
	}
}
