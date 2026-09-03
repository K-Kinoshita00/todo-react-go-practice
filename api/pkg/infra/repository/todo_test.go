package repository

import (
	"context"
	"testing"

	"github.com/google/uuid"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/dto"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/domain/entity"
)

func TestTodoRepositoryInsert(t *testing.T) {
	db := dbSetup(t)
	defer db.Close() // 終了後にDBを閉じる
	ctx := context.Background()
	cmd := NewTodoRepository(db)

	testTodoId := uuid.MustParse("01a06617-04db-72db-a2d7-894718bc83df")
	testTodoTitle := "test_todo"
	testTodoStatus := entity.TodoStatusNotStarted

	err := cmd.Insert(ctx, &entity.Todo{
		ID:     testTodoId,
		Title:  testTodoTitle,
		Status: testTodoStatus,
	})
	if err != nil {
		t.Fatalf("Insert: %v", err)
	}

	query := NewTodoQueryRepository(db)
	todo, err := query.FindByID(ctx, testTodoId)
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

func TestTodoRepositoryUpdate(t *testing.T) {
	db := dbSetup(t)
	defer db.Close() // 終了後にDBを閉じる
	ctx := context.Background()
	cmd := NewTodoRepository(db)

	testTodoId := uuid.MustParse("01a06617-04db-72db-a2d7-894718bc83df")
	testTodoTitle := "test_todo"
	testTodoStatus := entity.TodoStatusNotStarted

	err := cmd.Insert(ctx, &entity.Todo{
		ID:     testTodoId,
		Title:  testTodoTitle,
		Status: testTodoStatus,
	})
	if err != nil {
		t.Fatalf("Insert: %v", err)
	}
	err = cmd.Update(ctx, &entity.Todo{
		ID:     testTodoId,
		Title:  testTodoTitle,
		Status: entity.TodoStatusInProgress,
	})
	if err != nil {
		t.Fatalf("Update: %v", err)
	}

	query := NewTodoQueryRepository(db)
	todo, err := query.FindByID(ctx, testTodoId)
	if err != nil {
		t.Fatalf("FindByID: %v", err)
	}
	if todo.Title != testTodoTitle {
		t.Fatalf("Title: got %v, want %v", todo.Title, testTodoTitle)
	}
	if todo.Status != dto.TodoStatusInProgress {
		t.Fatalf("Status: got %v, want %v", todo.Status, dto.TodoStatusInProgress)
	}
}

func TestTodoRepositoryDelete(t *testing.T) {
	db := dbSetup(t)
	defer db.Close()
	ctx := context.Background()
	cmd := NewTodoRepository(db)

	testTodoId := uuid.MustParse("01a06617-04db-72db-a2d7-894718bc83df")
	testTodoTitle := "test_todo"
	testTodoStatus := entity.TodoStatusNotStarted

	err := cmd.Insert(ctx, &entity.Todo{
		ID:     testTodoId,
		Title:  testTodoTitle,
		Status: testTodoStatus,
	})
	if err != nil {
		t.Fatalf("Insert: %v", err)
	}
	err = cmd.Delete(ctx, testTodoId)
	if err != nil {
		t.Fatalf("Delete: %v", err)
	}

	query := NewTodoQueryRepository(db)
	todos, err := query.List(ctx)
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	var found dto.Todo
	for _, todo := range todos {
		if todo.ID == testTodoId {
			found = *todo
			break
		}
	}
	if found.ID == testTodoId {
		t.Fatalf("Delete: %v is not deleted", testTodoId)
	}
}
