package usecase

import (
	"context"

	"github.com/google/uuid"
)

type TodoStatus string

const (
	TodoStatusNotStarted TodoStatus = "not_started"
	TodoStatusInProgress TodoStatus = "in_progress"
	TodoStatusCompleted TodoStatus = "completed"
	TodoStatusArchive TodoStatus = "archive"
)

type InsertTodo struct {
	Title  string
	Status TodoStatus
}

type UpdateTodo struct {
	ID     uuid.UUID
	Title  string
	Status TodoStatus
}

type TodoCommandRepository interface {
	Insert(ctx context.Context, params InsertTodo) (uuid.UUID,error)
	Update(ctx context.Context, params UpdateTodo) error
	Delete(ctx context.Context, id uuid.UUID) error
}
