package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/dto"
	appErr "github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/error"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/domain/entity"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/interface/gen/openapi"
	middleware "github.com/K-Kinoshita00/todo-react-go-practice/pkg/interface/middleware"
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
	claims, ok := middleware.ClaimsFromContext(ctx)
	if !ok {
		appError := presenter.MapToAppError(ctx, appErr.ErrUnauthorized)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}

	todos, err := h.uc.List(ctx, claims.Sub)
	if err != nil {
		appError := presenter.MapToAppError(ctx, err)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
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
	claims, ok := middleware.ClaimsFromContext(ctx)
	if !ok {
		appError := presenter.MapToAppError(ctx, appErr.ErrUnauthorized)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}

	todo, err := h.uc.FindByID(ctx, id, claims.Sub)
	if err != nil {
		appError := presenter.MapToAppError(ctx, err)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
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
	claims, ok := middleware.ClaimsFromContext(ctx)
	if !ok {
		appError := presenter.MapToAppError(ctx, appErr.ErrUnauthorized)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}

	var req openapi.CreateTodo
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		appError := presenter.MapToAppError(ctx, fmt.Errorf("create todo json decode: %v: %w", err, appErr.ErrBadRequest))
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}

	err := h.uc.Create(ctx, req.Title, entity.TodoStatus(req.Status), claims.Sub)
	if err != nil {
		appError := presenter.MapToAppError(ctx, err)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}
	presenter.NewResponse(http.StatusCreated, nil).Send(w)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request, id openapi.ID) {
	ctx := r.Context()
	claims, ok := middleware.ClaimsFromContext(ctx)
	if !ok {
		appError := presenter.MapToAppError(ctx, appErr.ErrUnauthorized)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}

	var req openapi.UpdateTodo
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		appError := presenter.MapToAppError(ctx, fmt.Errorf("update todo json decode: %v: %w", err, appErr.ErrBadRequest))
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}

	err := h.uc.Update(ctx, id, req.Title, entity.TodoStatus(req.Status), claims.Sub)
	if err != nil {
		appError := presenter.MapToAppError(ctx, err)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}
	presenter.NewResponse(http.StatusNoContent, nil).Send(w)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request, id openapi.ID) {
	ctx := r.Context()
	claims, ok := middleware.ClaimsFromContext(ctx)
	if !ok {
		appError := presenter.MapToAppError(ctx, appErr.ErrUnauthorized)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}

	err := h.uc.Delete(ctx, id, claims.Sub)
	if err != nil {
		appError := presenter.MapToAppError(ctx, err)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}
	presenter.NewResponse(http.StatusNoContent, nil).Send(w)
}

type TodoUseCase interface {
	Create(ctx context.Context, title string, status entity.TodoStatus, owner string) error
	Update(ctx context.Context, id uuid.UUID, title string, status entity.TodoStatus, owner string) error
	Delete(ctx context.Context, id uuid.UUID, owner string) error
	List(ctx context.Context, owner string) ([]*dto.Todo, error)
	FindByID(ctx context.Context, id uuid.UUID, owner string) (*dto.Todo, error)
}
