package handler

import (
	"net/http"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/interface/presenter"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	presenter.NewResponse(http.StatusOK, map[string]string{"status": "ok"}).Send(w)
}
