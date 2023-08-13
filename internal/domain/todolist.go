package domain

import (
	"errors"
	"time"
	"unicode/utf8"
)

type Task struct {
	Title    string    `json:"title"`
	ActiveAt time.Time `json:"activeAt"`
	HashKey  string    `json:"hash"`
}

func (l *Task) Validate() error {
	if l.Title == "" {
		return errors.New("Title must not be empty")
	}
	if utf8.RuneCountInString(l.Title) > 200 {
		return errors.New("Title length must be maximum 200 symbols")
	}
	return nil
}
