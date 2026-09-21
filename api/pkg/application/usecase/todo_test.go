package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/dto"
	appErr "github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/error"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/domain/entity"
)

type fakeTodoRepository struct {
	todos        map[uuid.UUID]*entity.Todo
	insertCalled bool
	updateCalled bool
	deleteCalled bool
}

func (f *fakeTodoRepository) Insert(ctx context.Context, params *entity.Todo) error {
	if f.todos == nil {
		f.todos = make(map[uuid.UUID]*entity.Todo)
	}
	f.todos[params.ID] = params
	f.insertCalled = true
	return nil
}
func (f *fakeTodoRepository) Update(ctx context.Context, params *entity.Todo) error {
	if f.todos == nil {
		f.todos = make(map[uuid.UUID]*entity.Todo)
	}
	f.todos[params.ID] = params
	f.updateCalled = true
	return nil
}
func (f *fakeTodoRepository) Delete(ctx context.Context, id uuid.UUID, owner string) error {
	if f.todos == nil {
		f.todos = make(map[uuid.UUID]*entity.Todo)
	}
	delete(f.todos, id) // go 組み込みの関数 mapから特定のkeyを削除
	f.deleteCalled = true
	return nil
}
func (f *fakeTodoRepository) FindByID(ctx context.Context, id uuid.UUID, owner string) (*entity.Todo, error) {
	if f.todos == nil {
		f.todos = make(map[uuid.UUID]*entity.Todo)
	}
	if todo, ok := f.todos[id]; ok {
		if todo.Owner != owner {
			return nil, appErr.ErrNotFound
		}
		return todo, nil
	}
	return nil, nil
}

type fakeTodoQueryService struct {
	todos map[uuid.UUID]*dto.Todo
}

func (f *fakeTodoQueryService) List(ctx context.Context, owner string) ([]*dto.Todo, error) {
	if f.todos == nil {
		f.todos = make(map[uuid.UUID]*dto.Todo)
	}
	todos := make([]*dto.Todo, 0, len(f.todos))
	for _, todo := range f.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}
func (f *fakeTodoQueryService) FindByID(ctx context.Context, id uuid.UUID, owner string) (*dto.Todo, error) {
	if f.todos == nil {
		f.todos = make(map[uuid.UUID]*dto.Todo)
	}
	if todo, ok := f.todos[id]; ok {
		return todo, nil
	}
	return nil, nil
}

func TestCreateTodo_OK(t *testing.T) {
	fake := &fakeTodoRepository{}
	uc := NewTodoUseCase(fake, &fakeTodoQueryService{})
	err := uc.Create(t.Context(), "test_todo", entity.TodoStatusNotStarted, "test_user")
	if err != nil {
		t.Fatalf("CreateTodo_OK: %v", err)
	}
	if !fake.insertCalled {
		t.Fatalf("CreateTodo_OK: insertCalled is not true")
	}
}

func TestCreateTodo_BadRequest(t *testing.T) {
	fake := &fakeTodoRepository{}
	uc := NewTodoUseCase(fake, &fakeTodoQueryService{})
	err := uc.Create(t.Context(), "", entity.TodoStatusNotStarted, "test_user")
	if !errors.Is(err, appErr.ErrBadRequest) {
		t.Fatalf("CreateTodo_BadRequest empty title: %v", err)
	}
	if fake.insertCalled {
		t.Fatalf("CreateTodo_BadRequest: insertCalled is not false")
	}

	err = uc.Create(t.Context(), "test_todo", "", "test_user")
	if !errors.Is(err, appErr.ErrBadRequest) {
		t.Fatalf("CreateTodo_BadRequest bad status: %v", err)
	}
	if fake.insertCalled {
		t.Fatalf("CreateTodo_BadRequest: insertCalled is not false")
	}
}

func TestUpdateTodo_OK(t *testing.T) {
	id := uuid.New()
	fake := &fakeTodoRepository{
		todos: map[uuid.UUID]*entity.Todo{
			id: {
				ID:     id,
				Title:  "test_todo",
				Status: entity.TodoStatusNotStarted,
				Owner:  "test_user",
			},
		},
	}
	uc := NewTodoUseCase(fake, &fakeTodoQueryService{})
	err := uc.Update(t.Context(), id, "test_todo", entity.TodoStatusNotStarted, "test_user")
	if err != nil {
		t.Fatalf("UpdateTodo_OK: %v", err)
	}
	if !fake.updateCalled {
		t.Fatalf("UpdateTodo_OK: updateCalled is not true")
	}
}

func TestUpdateTodo_NotFound(t *testing.T) {
	id := uuid.New()
	fake := &fakeTodoRepository{
		todos: map[uuid.UUID]*entity.Todo{
			id: {
				ID:     id,
				Title:  "test_todo",
				Status: entity.TodoStatusNotStarted,
				Owner:  "test_user",
			},
		},
	}
	uc := NewTodoUseCase(fake, &fakeTodoQueryService{})
	err := uc.Update(t.Context(), uuid.New(), "test_todo", entity.TodoStatusNotStarted, "test_user")
	if !errors.Is(err, appErr.ErrNotFound) {
		t.Fatalf("UpdateTodo_NotFound: %v", err)
	}
	if fake.updateCalled {
		t.Fatalf("UpdateTodo_NotFound: updateCalled is not false")
	}

	// 別のユーザーが更新しようとした場合
	err = uc.Update(t.Context(), id, "test_todo", entity.TodoStatusNotStarted, "test_user2")
	if !errors.Is(err, appErr.ErrNotFound) {
		t.Fatalf("UpdateTodo_NotFound: wrong user %v", err)
	}
	if fake.updateCalled {
		t.Fatalf("UpdateTodo_NotFound: updateCalled is not false")
	}
}

func TestUpdateTodo_BadRequest(t *testing.T) {
	id := uuid.New()
	fake := &fakeTodoRepository{
		todos: map[uuid.UUID]*entity.Todo{
			id: {
				ID:     id,
				Title:  "test_todo",
				Status: entity.TodoStatusNotStarted,
				Owner:  "test_user",
			},
		},
	}
	uc := NewTodoUseCase(fake, &fakeTodoQueryService{})
	err := uc.Update(t.Context(), id, "", entity.TodoStatusNotStarted, "test_user")
	if !errors.Is(err, appErr.ErrBadRequest) {
		t.Fatalf("UpdateTodo_BadRequest empty title: %v", err)
	}
	if fake.updateCalled {
		t.Fatalf("UpdateTodo_BadRequest: updateCalled is not false")
	}

	err = uc.Update(t.Context(), id, "test_todo", "", "test_user")
	if !errors.Is(err, appErr.ErrBadRequest) {
		t.Fatalf("UpdateTodo_BadRequest bad status: %v", err)
	}
	if fake.updateCalled {
		t.Fatalf("UpdateTodo_BadRequest: updateCalled is not false")
	}
}

func TestDeleteTodo_OK(t *testing.T) {
	id := uuid.New()
	fake := &fakeTodoRepository{
		todos: map[uuid.UUID]*entity.Todo{
			id: {
				ID:     id,
				Title:  "test_todo",
				Status: entity.TodoStatusNotStarted,
				Owner:  "test_user",
			},
		},
	}
	uc := NewTodoUseCase(fake, &fakeTodoQueryService{})
	err := uc.Delete(t.Context(), id, "test_user")
	if err != nil {
		t.Fatalf("DeleteTodo_OK: %v", err)
	}
}

func TestDeleteTodo_NotFound(t *testing.T) {
	id := uuid.New()
	fake := &fakeTodoRepository{
		todos: map[uuid.UUID]*entity.Todo{
			id: {
				ID:     id,
				Title:  "test_todo",
				Status: entity.TodoStatusNotStarted,
				Owner:  "test_user",
			},
		},
	}
	uc := NewTodoUseCase(fake, &fakeTodoQueryService{})
	err := uc.Delete(t.Context(), uuid.New(), "test_user")
	if !errors.Is(err, appErr.ErrNotFound) {
		t.Fatalf("DeleteTodo_NotFound: %v", err)
	}
	if fake.deleteCalled != false {
		t.Fatalf("DeleteTodo_NotFound: deleteCalled is not false")
	}

	// 別のユーザーが削除しようとした場合
	err = uc.Delete(t.Context(), id, "test_user2")
	if !errors.Is(err, appErr.ErrNotFound) {
		t.Fatalf("DeleteTodo_NotFound wrong user: %v", err)
	}
	if fake.deleteCalled {
		t.Fatalf("DeleteTodo_NotFound: deleteCalled is not false")
	}
}

func TestListTodo_OK(t *testing.T) {
	id := uuid.New()
	fake := &fakeTodoQueryService{
		todos: map[uuid.UUID]*dto.Todo{
			id: {
				ID:     id,
				Title:  "test_todo",
				Status: dto.TodoStatusNotStarted,
			},
		},
	}
	uc := NewTodoUseCase(&fakeTodoRepository{}, fake)
	todos, err := uc.List(t.Context(), "test_user")
	if err != nil {
		t.Fatalf("ListTodo_OK: %v", err)
	}
	if len(todos) == 0 {
		t.Fatalf("ListTodo_OK: todos is empty")
	}
}

func TestFindByIDTodo_OK(t *testing.T) {
	id := uuid.New()
	fake := &fakeTodoQueryService{
		todos: map[uuid.UUID]*dto.Todo{
			id: {
				ID:     id,
				Title:  "test_todo",
				Status: dto.TodoStatusNotStarted,
			},
		},
	}
	uc := NewTodoUseCase(&fakeTodoRepository{}, fake)
	todo, err := uc.FindByID(t.Context(), id, "test_user")
	if err != nil {
		t.Fatalf("FindByIDTodo_OK: %v", err)
	}
	if todo == nil {
		t.Fatalf("FindByIDTodo_OK: todo is nil")
	}
}

func TestFindByIDTodo_NotFound(t *testing.T) {
	id := uuid.New()
	fake := &fakeTodoQueryService{
		todos: map[uuid.UUID]*dto.Todo{
			id: {
				ID:     id,
				Title:  "test_todo",
				Status: dto.TodoStatusNotStarted,
			},
		},
	}
	uc := NewTodoUseCase(&fakeTodoRepository{}, fake)
	todo, err := uc.FindByID(t.Context(), uuid.New(), "test_user")
	if !errors.Is(err, appErr.ErrNotFound) {
		t.Fatalf("FindByIDTodo_NotFound: %v", err)
	}
	if todo != nil {
		t.Fatalf("FindByIDTodo_NotFound: todo is not nil")
	}
}
