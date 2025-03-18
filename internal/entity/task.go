package entity

import "github.com/google/uuid"

type Task struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title" binding:"required"`
	Desc     string    `json:"desc"`
	Status   bool      `json:"status"`
	Priority string    `json:"priority"`
}

var Tasklist = map[uuid.UUID]Task{}
