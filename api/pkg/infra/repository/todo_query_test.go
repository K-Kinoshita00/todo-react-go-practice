package repository

import (
	"context"
	"testing"

	"github.com/google/uuid"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/dto"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/domain/entity"
)

func TestTodoQueryRepositoryFindByID(t *testing.T) {
	db := dbSetup(t)
	ctx := context.Background()
	query := NewTodoQueryRepository(db)
	cmd := NewTodoRepository(db)

	testTodoId := uuid.MustParse("01a06617-04db-72db-a2d7-894718bc83df")
	testTodoOwner := "01a0b3be-5066-701b-8304-1ec0436ac032"
	testTodoTitle := "test_todo"
	testCmdTodoStatus := entity.TodoStatusNotStarted

	err := cmd.Insert(ctx, &entity.Todo{
		ID:     testTodoId,
		Title:  testTodoTitle,
		Status: testCmdTodoStatus,
		Owner:  testTodoOwner,
	})
	if err != nil {
		t.Fatalf("Insert: %v", err)
	}
	t.Cleanup(func() {
		cmd.Delete(ctx, testTodoId, testTodoOwner)
		db.Close()
	})

	todo, err := query.FindByID(ctx, testTodoId, testTodoOwner)
	if err != nil {
		t.Fatalf("FindByID: %v", err)
	}

	if todo.Title != testTodoTitle {
		t.Fatalf("Title: got %v, want %v", todo.Title, testTodoTitle)
	}
	if todo.Status != dto.TodoStatusNotStarted {
		t.Fatalf("Status: got %v, want %v", todo.Status, dto.TodoStatusNotStarted)
	}
}

func TestTodoQueryRepositoryList(t *testing.T) {
	db := dbSetup(t)
	t.Cleanup(func() {
		db.Close()
	})
	ctx := context.Background()
	query := NewTodoQueryRepository(db)
	cmd := NewTodoRepository(db)

	testTodoOwner := "01a0b3be-5066-701b-8304-1ec0436ac032"
	TodoData1 := entity.Todo{
		ID:     uuid.MustParse("01a0661e-8efc-7459-8323-6911fb746c92"),
		Title:  "test_todo_1",
		Status: entity.TodoStatusNotStarted,
		Owner:  testTodoOwner,
	}
	TodoData2 := entity.Todo{
		ID:     uuid.MustParse("01a0661e-d896-7408-909c-62df844936f4"),
		Title:  "test_todo_2",
		Status: entity.TodoStatusInProgress,
		Owner:  testTodoOwner,
	}
	err := cmd.Insert(ctx, &TodoData1)
	if err != nil {
		t.Fatalf("Insert: %v", err)
	}
	t.Cleanup(func() {
		cmd.Delete(ctx, TodoData1.ID, testTodoOwner)
	})

	err = cmd.Insert(ctx, &TodoData2)
	if err != nil {
		t.Fatalf("Insert: %v", err)
	}
	t.Cleanup(func() {
		cmd.Delete(ctx, TodoData2.ID, testTodoOwner)
	})

	todos, err := query.List(ctx, testTodoOwner)
	if err != nil {
		t.Fatalf("List: %v", err)
	}

	var found1, found2 dto.Todo
	for _, todo := range todos {
		if todo.ID == TodoData1.ID {
			found1 = *todo
			if todo.Title != TodoData1.Title {
				t.Fatalf("Title: got %v, want %v", todo.Title, TodoData1.Title)
			}
			if todo.Status != dto.TodoStatusNotStarted {
				t.Fatalf("Status: got %v, want %v", todo.Status, dto.TodoStatusNotStarted)
			}
		}
		if todo.ID == TodoData2.ID {
			found2 = *todo
			if todo.Title != TodoData2.Title {
				t.Fatalf("Title: got %v, want %v", todo.Title, TodoData2.Title)
			}
			if todo.Status != dto.TodoStatusInProgress {
				t.Fatalf("Status: got %v, want %v", todo.Status, dto.TodoStatusInProgress)
			}
		}
	}
	if found1.ID != TodoData1.ID {
		t.Fatalf("found1.ID: got %v, want %v", found1.ID, TodoData1.ID)
	}
	if found2.ID != TodoData2.ID {
		t.Fatalf("found2.ID: got %v, want %v", found2.ID, TodoData2.ID)
	}
}
