package entity

import (
	"errors"

	"github.com/google/uuid"
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

func NewTodo(
	title string,
	status TodoStatus,
) (*Todo, error) {
	return &Todo{
		ID:     uuid.New(),
		Title:  title,
		Status: status,
	}, nil
}

func (e *Todo) UpdateTitleAndStatus(
	title string,
	status TodoStatus,
) error {
	e.Title = title
	e.Status = status
	return nil
}

func (e *Todo) Validate() error {
	if e.Title == "" {
		return errors.New("title is required")
	}
	if e.Status == "" {
		return errors.New("status is required")
	}
	return nil
}
