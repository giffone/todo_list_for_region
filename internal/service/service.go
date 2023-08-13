package service

import "todolist/internal/repository"

type service struct {
	db repository.Db
}

