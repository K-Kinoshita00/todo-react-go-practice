package repository

import (
	"context"
	"testing"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/queryservice"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/usecase"
)

func TestTodoCommandRepositoryInsert(t *testing.T) {
	db := dbSetup(t)
	defer db.Close() // 終了後にDBを閉じる
	ctx := context.Background()
	TestTodoTitle := "test_todo"
	TestTodoStatus := usecase.TodoStatusNotStarted
	cmd := NewTodoCommandRepository(db)
	id, err := cmd.Insert(ctx, usecase.InsertTodo{
		Title:  TestTodoTitle,
		Status: TestTodoStatus,
	})
	if err != nil {
		t.Fatalf("Insert: %v", err)
	}

	query := NewTodoQueryRepository(db)
	todo, err := query.FindByID(ctx, id)
	if err != nil {
		t.Fatalf("FindByID: %v", err)
	}
	if todo.Title != TestTodoTitle {
		t.Fatalf("Title: got %v, want %v", todo.Title, TestTodoTitle)
	}
	if todo.Status != queryservice.TodoStatusNotStarted {
		t.Fatalf("Status: got %v, want %v", todo.Status, queryservice.TodoStatusNotStarted)
	}
}

func TestTodoCommandRepositoryUpdate(t *testing.T) {
	db := dbSetup(t)
	defer db.Close() // 終了後にDBを閉じる
	ctx := context.Background()
	cmd := NewTodoCommandRepository(db)
	id, err := cmd.Insert(ctx, usecase.InsertTodo{
		Title:  "test_todo",
		Status: usecase.TodoStatusNotStarted,
	})
	if err != nil {
		t.Fatalf("Insert: %v", err)
	}
	TestTodoTitle := "test_todo_updated"
	err = cmd.Update(ctx, usecase.UpdateTodo{
		ID:     id,
		Title:  TestTodoTitle,
		Status: usecase.TodoStatusInProgress,
	})
	if err != nil {
		t.Fatalf("Update: %v", err)
	}

	query := NewTodoQueryRepository(db)
	todo, err := query.FindByID(ctx, id)
	if err != nil {
		t.Fatalf("FindByID: %v", err)
	}
	if todo.Title != TestTodoTitle {
		t.Fatalf("Title: got %v, want %v", todo.Title, TestTodoTitle)
	}
	if todo.Status != queryservice.TodoStatusInProgress {
		t.Fatalf("Status: got %v, want %v", todo.Status, queryservice.TodoStatusInProgress)
	}
}

func TestTodoCommandRepositoryDelete(t *testing.T) {
	db := dbSetup(t)
	defer db.Close()
	ctx := context.Background()
	cmd := NewTodoCommandRepository(db)
	id, err := cmd.Insert(ctx, usecase.InsertTodo{
		Title:  "test_todo",
		Status: usecase.TodoStatusNotStarted,
	})
	if err != nil {
		t.Fatalf("Insert: %v", err)
	}
	err = cmd.Delete(ctx, id)
	if err != nil {
		t.Fatalf("Delete: %v", err)
	}

	query := NewTodoQueryRepository(db)
	todos, err := query.List(ctx)
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	var found queryservice.Todo
	for _, todo := range todos {
		if todo.ID == id {
			found = *todo
			break
		}
	}
	if found.ID == id {
		t.Fatalf("Delete: %v is not deleted", id)
	}
}
