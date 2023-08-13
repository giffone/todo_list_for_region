package domain

import (
	"errors"
	"time"
	"unicode/utf8"
)

type Request struct {
	Title     string `json:"title"`
	ActiveAt  string `json:"activeAt"`
	ValidDate time.Time
}

func (r *Request) Validate() error {
	var err error
	if r.Title == "" {
		return errors.New("Title must not be empty")
	}
	if utf8.RuneCountInString(r.Title) > 200 {
		return errors.New("Title length must be maximum 200 symbols")
	}
	r.ValidDate, err = time.Parse(time.RFC3339, string(r.ActiveAt))
	if err != nil {
		return errors.New("Not valid date format, must be as yyyy-mm-dd")
	}
	return nil
}

type TaskDTO struct {
	Title    string    `bson:"title"`
	ActiveAt time.Time `bson:"activeAt"`
	HashKey  string    `bson:"hash"`
}
