package domain

import (
	"time"
)

type Status int64

const (
	CREATED uint = iota
	CREATION_ERROR
	EXECUTED
	EXECUTION_ERROR
)

type Task struct {
	ID         uint32
	CreatedAt  time.Time
	FinishedAt time.Time
	Status     Status
}
