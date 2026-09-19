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
	Owner  string
}

func NewTodo(
	title string,
	status TodoStatus,
	owner string,
) (*Todo, error) {
	return &Todo{
		ID:     uuid.New(),
		Title:  title,
		Status: status,
		Owner:  owner,
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
	if e.Status != TodoStatusNotStarted && e.Status != TodoStatusInProgress && e.Status != TodoStatusCompleted && e.Status != TodoStatusArchive {
		return errors.New("status is invalid")
	}
	if e.Owner == "" {
		return errors.New("owner is required")
	}
	return nil
}
