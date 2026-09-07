package presenter

import (
	"context"

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
