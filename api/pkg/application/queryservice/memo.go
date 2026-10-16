package queryservice

import (
	"context"
	"uuid"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/dto"
)

type MemoQueryService interface {
	List(ctx context.Context, owner string) ([]*dto.Memo, error)
	FindByID(ctx context.Context, id uuid.UUID, owner string) (*dto.Memo, error)
}
