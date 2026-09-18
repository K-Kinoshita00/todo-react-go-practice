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

	testTodoId := uuid.MustParse("01a0b3df-50ff-711c-ba87-56fdcc3d2e4f")
	testTodoOwner := "01a0b3be-5066-701b-8304-1ec0436ac032"
	testTodoTitle := "test_todo"
	testTodoStatus := entity.TodoStatusNotStarted

	err := cmd.Insert(ctx, &entity.Todo{
		ID:     testTodoId,
		Title:  testTodoTitle,
		Status: testTodoStatus,
		Owner:  testTodoOwner,
	})
	if err != nil {
		t.Fatalf("Insert: %v", err)
	}

	query := NewTodoQueryRepository(db)
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

func TestTodoRepositoryUpdate(t *testing.T) {
	db := dbSetup(t)
	defer db.Close() // 終了後にDBを閉じる
	ctx := context.Background()
	cmd := NewTodoRepository(db)

	testTodoId := uuid.MustParse("01a0b3df-61e5-752c-ad6e-bbf9847ef764")
	testTodoOwner := "01a0b3be-5066-701b-8304-1ec0436ac032"
	testTodoTitle := "test_todo"
	testTodoStatus := entity.TodoStatusNotStarted

	err := cmd.Insert(ctx, &entity.Todo{
		ID:     testTodoId,
		Title:  testTodoTitle,
		Status: testTodoStatus,
		Owner:  testTodoOwner,
	})
	if err != nil {
		t.Fatalf("Insert: %v", err)
	}
	err = cmd.Update(ctx, &entity.Todo{
		ID:     testTodoId,
		Title:  testTodoTitle,
		Status: entity.TodoStatusInProgress,
		Owner:  testTodoOwner,
	})
	if err != nil {
		t.Fatalf("Update: %v", err)
	}

	query := NewTodoQueryRepository(db)
	todo, err := query.FindByID(ctx, testTodoId, testTodoOwner)
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

	testTodoId := uuid.MustParse("01a0b3df-7492-73eb-9d59-447d4e6f85a4")
	testTodoOwner := "01a0b3be-5066-701b-8304-1ec0436ac032"
	testTodoTitle := "test_todo"
	testTodoStatus := entity.TodoStatusNotStarted

	err := cmd.Insert(ctx, &entity.Todo{
		ID:     testTodoId,
		Title:  testTodoTitle,
		Status: testTodoStatus,
		Owner:  testTodoOwner,
	})
	if err != nil {
		t.Fatalf("Insert: %v", err)
	}
	err = cmd.Delete(ctx, testTodoId, testTodoOwner)
	if err != nil {
		t.Fatalf("Delete: %v", err)
	}

	query := NewTodoQueryRepository(db)
	todos, err := query.List(ctx, testTodoOwner)
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
