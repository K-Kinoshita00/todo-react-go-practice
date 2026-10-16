package dto

import (
	"uuid"
)

type TodoStatus string

const (
	TodoStatusNotStarted TodoStatus = "not_started"
	TodoStatusInProgress TodoStatus = "in_progress"
	TodoStatusCompleted  TodoStatus = "completed"
	TodoStatusArchive    TodoStatus = "archive"
)

type Todo struct {
	ID     uuid.UUID
	Title  string
	Status TodoStatus
}
