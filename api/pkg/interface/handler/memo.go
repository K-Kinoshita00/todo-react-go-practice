package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"uuid"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/dto"
	appErr "github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/error"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/interface/gen/openapi"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/interface/middleware"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/interface/presenter"
)

// 外部が知る必要ないため小文字
type MemoHandler struct {
	uc MemoUseCase
}

func NewMemoHandler(uc MemoUseCase) *MemoHandler {
	return &MemoHandler{uc: uc}
}

func (h *MemoHandler) ListMemos(w http.ResponseWriter, r *http.Request, params openapi.ListMemosParams) {
	ctx := r.Context()
	claims, ok := middleware.ClaimsFromContext(ctx)
	if !ok {
		appError := presenter.MapToAppError(ctx, appErr.ErrUnauthorized)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}

	memos, err := h.uc.List(ctx, claims.Sub)
	if err != nil {
		appError := presenter.MapToAppError(ctx, err)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}

	resMemos := []openapi.Memo{}
	for _, memo := range memos {
		resMemos = append(resMemos, openapi.Memo{
			Id:      memo.ID,
			Title:   memo.Title,
			Content: memo.Content,
		})
	}

	body := openapi.MemoWithPagination{
		Data:       resMemos,
		PageNumber: 1,
		PageSize:   len(resMemos),
		Total:      len(resMemos),
	}
	presenter.NewResponse(http.StatusOK, body).Send(w)
}

func (h *MemoHandler) GetMemoByID(w http.ResponseWriter, r *http.Request, id openapi.ID) {
	ctx := r.Context()
	claims, ok := middleware.ClaimsFromContext(ctx)
	if !ok {
		appError := presenter.MapToAppError(ctx, appErr.ErrUnauthorized)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}

	memo, err := h.uc.FindByID(ctx, id, claims.Sub)
	if err != nil {
		appError := presenter.MapToAppError(ctx, err)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}

	body := openapi.Memo{
		Id:      memo.ID,
		Title:   memo.Title,
		Content: memo.Content,
	}
	presenter.NewResponse(http.StatusOK, body).Send(w)
}

func (h *MemoHandler) CreateMemo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	claims, ok := middleware.ClaimsFromContext(ctx)
	if !ok {
		appError := presenter.MapToAppError(ctx, appErr.ErrUnauthorized)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}

	var req openapi.CreateMemo
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		appError := presenter.MapToAppError(ctx, appErr.ErrBadRequest)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}

	if err := h.uc.Create(ctx, req.Title, req.Content, claims.Sub); err != nil {
		appError := presenter.MapToAppError(ctx, err)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}

	presenter.NewResponse(http.StatusCreated, nil).Send(w)
}

func (h *MemoHandler) UpdateMemo(w http.ResponseWriter, r *http.Request, id openapi.ID) {
	ctx := r.Context()
	claims, ok := middleware.ClaimsFromContext(ctx)
	if !ok {
		appError := presenter.MapToAppError(ctx, appErr.ErrUnauthorized)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}

	var req openapi.UpdateMemo
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		appError := presenter.MapToAppError(ctx, appErr.ErrBadRequest)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}

	if err := h.uc.Update(ctx, id, req.Title, req.Content, claims.Sub); err != nil {
		appError := presenter.MapToAppError(ctx, err)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}

	presenter.NewResponse(http.StatusOK, nil).Send(w)
}

func (h *MemoHandler) DeleteMemo(w http.ResponseWriter, r *http.Request, id openapi.ID) {
	ctx := r.Context()
	claims, ok := middleware.ClaimsFromContext(ctx)
	if !ok {
		appError := presenter.MapToAppError(ctx, appErr.ErrUnauthorized)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}

	if err := h.uc.Delete(ctx, id, claims.Sub); err != nil {
		appError := presenter.MapToAppError(ctx, err)
		presenter.NewResponse(appError.StatusCode, appError).Send(w)
		return
	}

	presenter.NewResponse(http.StatusNoContent, nil).Send(w)
}

type MemoUseCase interface {
	Create(ctx context.Context, title string, content string, owner string) error
	Update(ctx context.Context, id uuid.UUID, title string, content string, owner string) error
	Delete(ctx context.Context, id uuid.UUID, owner string) error
	List(ctx context.Context, owner string) ([]*dto.Memo, error)
	FindByID(ctx context.Context, id uuid.UUID, owner string) (*dto.Memo, error)
}
