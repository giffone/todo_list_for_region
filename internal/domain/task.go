package domain

import "time"

type TaskDTO struct {
	Title    string    `bson:"title"`
	ActiveAt time.Time `bson:"activeAt"`
	Status   bool      `bson:"status"`
	HashKey  string    `bson:"hash"`
}
