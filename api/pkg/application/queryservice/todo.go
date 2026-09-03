package queryservice

import (
	"context"

	"github.com/google/uuid"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/dto"
)

type TodoQueryService interface {
	List(ctx context.Context) ([]*dto.Todo, error)
	FindByID(ctx context.Context, id uuid.UUID) (*dto.Todo, error)
}
