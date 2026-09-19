package queryservice

import (
	"context"

	"github.com/google/uuid"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/dto"
)

type TodoQueryService interface {
	List(ctx context.Context, owner string) ([]*dto.Todo, error)
	FindByID(ctx context.Context, id uuid.UUID, owner string) (*dto.Todo, error)
}
