package usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/dto"
	appErr "github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/error"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/queryservice"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/domain/entity"
)

type TodoUseCase struct {
	cmd   TodoRepository
	query queryservice.TodoQueryService
}

func NewTodoUseCase(cmd TodoRepository, query queryservice.TodoQueryService) *TodoUseCase {
	return &TodoUseCase{cmd, query}
}

func (u *TodoUseCase) Create(ctx context.Context, title string, status entity.TodoStatus) error {
	ent, err := entity.NewTodo(title, status)
	if err != nil || ent == nil {
		return err
	}

	if res := ent.Validate(); res != nil {
		return appErr.ErrBadRequest
	}

	if err = u.cmd.Insert(ctx, ent); err != nil {
		return err
	}
	return nil
}

func (u *TodoUseCase) Update(ctx context.Context, id uuid.UUID, title string, status entity.TodoStatus) error {
	ent, err := u.cmd.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return appErr.ErrNotFound
		}
		return err
	}
	if ent == nil {
		return appErr.ErrNotFound
	}

	if err = ent.UpdateTitleAndStatus(title, status); err != nil {
		return appErr.ErrBadRequest
	}

	if err = u.cmd.Update(ctx, ent); err != nil {
		return err
	}
	return nil
}

func (u *TodoUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	ent, err := u.cmd.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return appErr.ErrNotFound
		}
		return err
	}
	if ent == nil {
		return appErr.ErrNotFound
	}

	if err = u.cmd.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (u *TodoUseCase) List(ctx context.Context) ([]*dto.Todo, error) {
	todos, err := u.query.List(ctx)
	if err != nil {
		return nil, err
	}
	return todos, err
}

func (u *TodoUseCase) FindByID(ctx context.Context, id uuid.UUID) (*dto.Todo, error) {
	todo, err := u.query.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, appErr.ErrNotFound
		}
		return nil, err
	}
	if todo == nil {
		return nil, appErr.ErrNotFound
	}
	return todo, nil
}

type TodoRepository interface {
	Insert(ctx context.Context, params *entity.Todo) error
	Update(ctx context.Context, params *entity.Todo) error
	Delete(ctx context.Context, id uuid.UUID) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Todo, error)
}
