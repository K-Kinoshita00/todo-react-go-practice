package presenter

import (
	"context"
	"net/http"

	appErr "github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/error"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/interface/gen/openapi"
)

func MapToAppError(ctx context.Context, err error) *openapi.Error {
	if err == nil {
		return nil
	}

	appError := appErr.NewAppError(err)
	return &openapi.Error{
		StatusCode: appError.StatusCode(),
		Code:       appError.Code(),
		Message:    appError.Message(),
	}
}

// パラメータの不一致による err の内容をクライアントに公開しないように変換
func BindError(w http.ResponseWriter, r *http.Request, err error) {
	appError := MapToAppError(r.Context(), appErr.ErrBadRequest)
	NewResponse(appError.StatusCode, appError).Send(w)
}
