package domain

import "time"

type TaskDTO struct {
	Title    string    `bson:"title"`
	ActiveAt time.Time `bson:"activeAt"`
	HashKey  string    `bson:"hash"`
}
