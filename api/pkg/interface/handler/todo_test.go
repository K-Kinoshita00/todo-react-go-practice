package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/dto"
	appErr "github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/error"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/domain/entity"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/interface/gen/openapi"
)

type fakeTodoUseCase struct {
	list []*dto.Todo
	err  error
}

func (f *fakeTodoUseCase) Create(ctx context.Context, title string, status entity.TodoStatus) error {
	return f.err
}

func (f *fakeTodoUseCase) Update(ctx context.Context, id uuid.UUID, title string, status entity.TodoStatus) error {
	return f.err
}

func (f *fakeTodoUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return f.err
}

func (f *fakeTodoUseCase) List(ctx context.Context) ([]*dto.Todo, error) {
	return f.list, f.err
}

func (f *fakeTodoUseCase) FindByID(ctx context.Context, id uuid.UUID) (*dto.Todo, error) {
	for _, todo := range f.list {
		if todo.ID == id {
			return todo, f.err
		}
	}
	return nil, f.err
}

func TestListTodos_OK(t *testing.T) {
	h := NewTodoHandler(&fakeTodoUseCase{
		list: []*dto.Todo{{
			ID:     uuid.MustParse("01a07b4c-257a-772e-aca3-f6abc0c941f8"),
			Title:  "test_todo1",
			Status: dto.TodoStatusNotStarted,
		}, {
			ID:     uuid.MustParse("01a07b4c-450b-74b8-a3f8-453d100baeaf"),
			Title:  "test_todo2",
			Status: dto.TodoStatusInProgress,
		}},
		err: nil,
	})
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec := httptest.NewRecorder()
	h.ListTodos(rec, req, openapi.ListTodosParams{})
	if rec.Code != http.StatusOK {
		t.Fatalf("status: got %d, want %d", rec.Code, http.StatusOK)
	}
	var body openapi.TodoWithPagination
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(body.Data) == 0 {
		t.Fatalf("length: got %d, want %d", len(body.Data), 1)
	}
	if body.Data[0].Title != "test_todo1" {
		t.Fatalf("title: got %s, want %s", body.Data[0].Title, "test_todo1")
	}
	if body.Data[1].Title != "test_todo2" {
		t.Fatalf("title: got %s, want %s", body.Data[1].Title, "test_todo2")
	}
}

func TestListTodos_BadRequest(t *testing.T) {
	testErr := appErr.ErrBadRequest
	h := NewTodoHandler(&fakeTodoUseCase{
		err: testErr,
	})
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec := httptest.NewRecorder()
	h.ListTodos(rec, req, openapi.ListTodosParams{})
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status: got %d, want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestGetTodoByID(t *testing.T) {
	testID := uuid.MustParse("01a07b4c-257a-772e-aca3-f6abc0c941f8")
	h := NewTodoHandler(&fakeTodoUseCase{
		list: []*dto.Todo{{
			ID:     testID,
			Title:  "test_todo1",
			Status: dto.TodoStatusNotStarted,
		}},
		err: nil,
	})
	req := httptest.NewRequest(http.MethodGet, "/todos/"+testID.String(), nil)
	rec := httptest.NewRecorder()
	h.GetTodoByID(rec, req, openapi.ID(testID))
	if rec.Code != http.StatusOK {
		t.Fatalf("status: got %d, want %d", rec.Code, http.StatusOK)
	}
	var todo dto.Todo
	json.NewDecoder(rec.Body).Decode(&todo)
	if todo.Title != "test_todo1" {
		t.Fatalf("title: got %s, want %s", todo.Title, "test_todo1")
	}
	if todo.Status != dto.TodoStatusNotStarted {
		t.Fatalf("status: got %s, want %s", todo.Status, dto.TodoStatusNotStarted)
	}
}

func TestGetTodoByID_NotFound(t *testing.T) {
	testID := uuid.MustParse("01a07b4c-257a-772e-aca3-f6abc0c941f8")
	h := NewTodoHandler(&fakeTodoUseCase{
		err: appErr.ErrNotFound,
	})
	req := httptest.NewRequest(http.MethodGet, "/todos/"+testID.String(), nil)
	rec := httptest.NewRecorder()
	h.GetTodoByID(rec, req, openapi.ID(testID))
	if rec.Code != http.StatusNotFound {
		t.Fatalf("status: got %d, want %d", rec.Code, http.StatusNotFound)
	}
}

func TestCreateTodo_OK(t *testing.T) {
	h := NewTodoHandler(&fakeTodoUseCase{
		err: nil,
	})
	body := openapi.CreateTodoJSONRequestBody{
		Title:  "test_todo1",
		Status: openapi.CreateTodoStatusNotStarted,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(bodyBytes))
	rec := httptest.NewRecorder()
	h.CreateTodo(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("status: got %d, want %d", rec.Code, http.StatusCreated)
	}
}

func TestCreateTodo_BadRequest(t *testing.T) {
	testID := uuid.MustParse("01a07b4c-257a-772e-aca3-f6abc0c941f8")
	testErr := appErr.ErrBadRequest
	h := NewTodoHandler(&fakeTodoUseCase{
		list: []*dto.Todo{{
			ID:     testID,
			Title:  "test_todo",
			Status: dto.TodoStatusNotStarted,
		}},
		err: testErr,
	})
	body := openapi.CreateTodoJSONRequestBody{
		Title:  "",
		Status: openapi.CreateTodoStatusNotStarted,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(bodyBytes))
	rec := httptest.NewRecorder()
	h.CreateTodo(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status: got %d, want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestUpdateTodo_OK(t *testing.T) {
	testID := uuid.MustParse("01a07b4c-257a-772e-aca3-f6abc0c941f8")
	h := NewTodoHandler(&fakeTodoUseCase{
		err: nil,
	})
	body := openapi.UpdateTodoJSONRequestBody{
		Title:  "test_todo",
		Status: openapi.UpdateTodoStatusNotStarted,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	req := httptest.NewRequest(http.MethodPut, "/todos/"+testID.String(), bytes.NewBuffer(bodyBytes))
	rec := httptest.NewRecorder()
	h.UpdateTodo(rec, req, openapi.ID(testID))
	if rec.Code != http.StatusNoContent {
		t.Fatalf("status: got %d, want %d", rec.Code, http.StatusNoContent)
	}
}

func TestUpdateTodo_BadRequest(t *testing.T) {
	testID := uuid.MustParse("01a07b4c-257a-772e-aca3-f6abc0c941f8")
	testErr := appErr.ErrBadRequest
	h := NewTodoHandler(&fakeTodoUseCase{
		list: []*dto.Todo{{
			ID:     testID,
			Title:  "test_todo",
			Status: dto.TodoStatusNotStarted,
		}},
		err: testErr,
	})
	body := openapi.UpdateTodoJSONRequestBody{
		Title:  "",
		Status: openapi.UpdateTodoStatusNotStarted,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	req := httptest.NewRequest(http.MethodPut, "/todos/"+testID.String(), bytes.NewBuffer(bodyBytes))
	rec := httptest.NewRecorder()
	h.UpdateTodo(rec, req, openapi.ID(testID))
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status: got %d, want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestUpdateTodo_NotFound(t *testing.T) {
	testID := uuid.MustParse("01a07b4c-257a-772e-aca3-f6abc0c941f8")
	h := NewTodoHandler(&fakeTodoUseCase{
		err: appErr.ErrNotFound,
	})
	body := openapi.UpdateTodoJSONRequestBody{
		Title:  "test_todo",
		Status: openapi.UpdateTodoStatusNotStarted,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	req := httptest.NewRequest(http.MethodPut, "/todos/"+testID.String(), bytes.NewBuffer(bodyBytes))
	rec := httptest.NewRecorder()
	h.UpdateTodo(rec, req, openapi.ID(testID))
	if rec.Code != http.StatusNotFound {
		t.Fatalf("status: got %d, want %d", rec.Code, http.StatusNotFound)
	}
}

func TestDeleteTodo_OK(t *testing.T) {
	testID := uuid.MustParse("01a07b4c-257a-772e-aca3-f6abc0c941f8")
	h := NewTodoHandler(&fakeTodoUseCase{
		err: nil,
	})
	req := httptest.NewRequest(http.MethodDelete, "/todos/"+testID.String(), nil)
	rec := httptest.NewRecorder()
	h.DeleteTodo(rec, req, openapi.ID(testID))
	if rec.Code != http.StatusNoContent {
		t.Fatalf("status: got %d, want %d", rec.Code, http.StatusNoContent)
	}
}

func TestDeleteTodo_NotFound(t *testing.T) {
	testID := uuid.MustParse("01a07b4c-257a-772e-aca3-f6abc0c941f8")
	h := NewTodoHandler(&fakeTodoUseCase{
		err: appErr.ErrNotFound,
	})
	req := httptest.NewRequest(http.MethodDelete, "/todos/"+testID.String(), nil)
	rec := httptest.NewRecorder()
	h.DeleteTodo(rec, req, openapi.ID(testID))
	if rec.Code != http.StatusNotFound {
		t.Fatalf("status: got %d, want %d", rec.Code, http.StatusNotFound)
	}
}
