package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/dto"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/domain/entity"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/interface/gen/openapi"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/interface/presenter"
)

// 外部が知る必要ないため小文字
type TodoHandler struct {
	uc TodoUseCase
}

func NewTodoHandler(uc TodoUseCase) *TodoHandler {
	return &TodoHandler{uc: uc}
}

func (h *TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request, params openapi.ListTodosParams) {
	ctx := r.Context()

	todos, err := h.uc.List(ctx)
	if err != nil {
		appErr := presenter.MapToAppError(ctx, err)
		presenter.NewResponse(appErr.StatusCode, appErr).Send(w)
		return
	}

	resTodos := []openapi.Todo{}

	for _, todo := range todos {
		resTodos = append(resTodos, openapi.Todo{
			Id:     todo.ID,
			Title:  todo.Title,
			Status: openapi.TodoStatus(todo.Status),
		})
	}

	body := openapi.TodoWithPagination{
		Data:       resTodos,
		PageNumber: 1,
		PageSize:   len(resTodos),
		Total:      len(resTodos),
	}
	presenter.NewResponse(http.StatusOK, body).Send(w)
}

func (h *TodoHandler) GetTodoByID(w http.ResponseWriter, r *http.Request, id openapi.ID) {
	ctx := r.Context()
	todo, err := h.uc.FindByID(ctx, id)
	if err != nil {
		appErr := presenter.MapToAppError(ctx, err)
		presenter.NewResponse(appErr.StatusCode, appErr).Send(w)
		return
	}

	res := openapi.TodoObjective{
		Id:     todo.ID,
		Title:  todo.Title,
		Status: openapi.TodoStatus(todo.Status),
	}
	presenter.NewResponse(http.StatusOK, res).Send(w)
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req openapi.CreateTodo
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		appErr := presenter.MapToAppError(ctx, err)
		presenter.NewResponse(appErr.StatusCode, appErr).Send(w)
		return
	}

	err := h.uc.Create(ctx, req.Title, entity.TodoStatus(req.Status))
	if err != nil {
		appErr := presenter.MapToAppError(ctx, err)
		presenter.NewResponse(appErr.StatusCode, appErr).Send(w)
		return
	}
	presenter.NewResponse(http.StatusNoContent, nil).Send(w)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request, id openapi.ID) {
	ctx := r.Context()
	var req openapi.UpdateTodo
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		appErr := presenter.MapToAppError(ctx, err)
		presenter.NewResponse(appErr.StatusCode, appErr).Send(w)
		return
	}

	err := h.uc.Update(ctx, id, req.Title, entity.TodoStatus(req.Status))
	if err != nil {
		appErr := presenter.MapToAppError(ctx, err)
		presenter.NewResponse(appErr.StatusCode, appErr).Send(w)
		return
	}
	presenter.NewResponse(http.StatusNoContent, nil).Send(w)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request, id openapi.ID) {
	ctx := r.Context()
	err := h.uc.Delete(ctx, id)
	if err != nil {
		appErr := presenter.MapToAppError(ctx, err)
		presenter.NewResponse(appErr.StatusCode, appErr).Send(w)
		return
	}
	presenter.NewResponse(http.StatusNoContent, nil).Send(w)
}

type TodoUseCase interface {
	Create(ctx context.Context, title string, status entity.TodoStatus) error
	Update(ctx context.Context, id uuid.UUID, title string, status entity.TodoStatus) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context) ([]*dto.Todo, error)
	FindByID(ctx context.Context, id uuid.UUID) (*dto.Todo, error)
}
