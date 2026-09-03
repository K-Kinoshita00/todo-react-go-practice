package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/dto"
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
	newTodoEnt, err := entity.NewTodo(title, status)
	if err != nil || newTodoEnt == nil {
		return err
	}

	err = u.cmd.Insert(ctx, newTodoEnt)
	return err
}

func (u *TodoUseCase) Update(ctx context.Context, id uuid.UUID, title string, status entity.TodoStatus) error {
	ent, err := u.cmd.FindByID(ctx, id)
	if err != nil || ent == nil {
		return err
	}

	err = ent.UpdateTitleAndStatus(title, status)
	if err != nil {
		return err
	}

	err = u.cmd.Update(ctx, ent)
	return err
}

func (u *TodoUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	ent, err := u.cmd.FindByID(ctx, id)
	if err != nil || ent == nil {
		return err
	}

	err = u.cmd.Delete(ctx, id)
	return err
}

func (u *TodoUseCase) List(ctx context.Context) ([]*dto.Todo, error) {
	todos, err := u.query.List(ctx)
	if err != nil {
		return nil, err
	}
	return todos, err
}

type TodoRepository interface {
	Insert(ctx context.Context, params *entity.Todo) error
	Update(ctx context.Context, params *entity.Todo) error
	Delete(ctx context.Context, id uuid.UUID) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Todo, error)
}
