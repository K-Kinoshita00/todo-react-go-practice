package handler

import (
	"encoding/json"
	"net/http"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/usecase"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/domain/entity"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/interface/gen/openapi"
)

type todoHandler struct {
	uc *usecase.TodoUseCase
}

func NewTodoHandler(uc *usecase.TodoUseCase) *todoHandler {
	return &todoHandler{uc: uc}
}

func (h *todoHandler) ListTodos(w http.ResponseWriter, r *http.Request, params openapi.ListTodosParams) {
	todos, err := h.uc.List(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func (h *todoHandler) GetTodoByID(w http.ResponseWriter, r *http.Request, id openapi.ID) {
	todo, err := h.uc.FindByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := openapi.TodoObjective{
		Id:     todo.ID,
		Title:  todo.Title,
		Status: openapi.TodoStatus(todo.Status),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *todoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req openapi.CreateTodo
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.uc.Create(r.Context(), req.Title, entity.TodoStatus(req.Status))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *todoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request, id openapi.ID) {
	var req openapi.UpdateTodo
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.uc.Update(r.Context(), id, req.Title, entity.TodoStatus(req.Status))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *todoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request, id openapi.ID) {
	err := h.uc.Delete(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *todoHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}