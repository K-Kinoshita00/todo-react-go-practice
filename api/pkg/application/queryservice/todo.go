package queryservice

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type TodoStatus string

const (
	TodoStatusNotStarted TodoStatus = "not_started"
	TodoStatusInProgress TodoStatus = "in_progress"
	TodoStatusCompleted TodoStatus = "completed"
	TodoStatusArchive TodoStatus = "archive"
)

type Todo struct {
	ID uuid.UUID
	Title string
	Status TodoStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TodoQueryService interface {
	List(ctx context.Context) ([]*Todo, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Todo, error)
}
