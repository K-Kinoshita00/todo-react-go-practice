package repository

import (
	"context"
	"testing"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/queryservice"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/usecase"
)

func TestTodoQueryRepositoryFindByID(t *testing.T) {
	db := dbSetup(t)
	defer db.Close()
	ctx := context.Background()
	query := NewTodoQueryRepository(db)
	cmd := NewTodoCommandRepository(db)

	TestTodoTitle := "test_todo"
	TestCmdTodoStatus := usecase.TodoStatusNotStarted

	id, err := cmd.Insert(ctx, usecase.InsertTodo{
		Title:  TestTodoTitle,
		Status: TestCmdTodoStatus,
	})
	if err != nil {
		t.Fatalf("Insert: %v", err)
	}
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

func TestTodoQueryRepositoryList(t *testing.T) {
	db := dbSetup(t)
	defer db.Close()
	ctx := context.Background()
	query := NewTodoQueryRepository(db)
	cmd := NewTodoCommandRepository(db)

	TodoData1 := usecase.InsertTodo{
		Title:  "test_todo_1",
		Status: usecase.TodoStatusNotStarted,
	}
	TodoData2 := usecase.InsertTodo{
		Title:  "test_todo_2",
		Status: usecase.TodoStatusInProgress,
	}
	id1, err := cmd.Insert(ctx, TodoData1)
	if err != nil {
		t.Fatalf("Insert: %v", err)
	}
	id2, err := cmd.Insert(ctx, TodoData2)
	if err != nil {
		t.Fatalf("Insert: %v", err)
	}

	todos, err := query.List(ctx)
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	var found1, found2 queryservice.Todo
	for _, todo := range todos {
		if todo.ID == id1 {
			found1 = *todo
			if todo.Title != TodoData1.Title {
				t.Fatalf("Title: got %v, want %v", todo.Title, TodoData1.Title)
			}
			if todo.Status != queryservice.TodoStatusNotStarted {
				t.Fatalf("Status: got %v, want %v", todo.Status, queryservice.TodoStatusNotStarted)
			}
		}
		if todo.ID == id2 {
			found2 = *todo
			if todo.Title != TodoData2.Title {
				t.Fatalf("Title: got %v, want %v", todo.Title, TodoData2.Title)
			}
			if todo.Status != queryservice.TodoStatusInProgress {
				t.Fatalf("Status: got %v, want %v", todo.Status, queryservice.TodoStatusInProgress)
			}
		}
	}
	if found1.ID != id1 {
		t.Fatalf("found1.ID: got %v, want %v", found1.ID, id1)
	}
	if found2.ID != id2 {
		t.Fatalf("found2.ID: got %v, want %v", found2.ID, id2)
	}
}
