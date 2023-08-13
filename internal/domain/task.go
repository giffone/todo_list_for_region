package domain

import "time"

type statuses string

func (s statuses) String() string { return string(s) }

const (
	StatusActive statuses = "active"
	StatusDone   statuses = "done"
)

type TaskDTO struct {
	Title    string    `bson:"title"`
	ActiveAt time.Time `bson:"activeAt"`
	Status   statuses  `bson:"status"`
	HashKey  string    `bson:"hash"`
}

type Task struct {
	ID       string    `bson:"id" json:"id"`
	Title    string    `bson:"title" json:"title"`
	ActiveAt time.Time `bson:"activeAt" json:"activeAt"`
}
