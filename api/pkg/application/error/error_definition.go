package error

import (
	"net/http"
)

type definedError struct {
	message    string
	statusCode int
}

var definedErrors = map[error]definedError{
	ErrBadRequest: {
		message:    "不正なリクエストです",
		statusCode: http.StatusBadRequest,
	},
	ErrUnauthorized: {
		message:    "認証が必要です",
		statusCode: http.StatusUnauthorized,
	},
	ErrBearerTokenExpired: {
		message:    "トークンが有効期限切れです",
		statusCode: http.StatusUnauthorized,
	},
	ErrForbidden: {
		message:    "アクセスが許可されていません",
		statusCode: http.StatusForbidden,
	},
	ErrNotFound: {
		message:    "リソースが見つかりません",
		statusCode: http.StatusNotFound,
	},
	ErrMethodNotAllowed: {
		message:    "許可されていないメソッドです",
		statusCode: http.StatusMethodNotAllowed,
	},
	ErrInternalServer: {
		message:    "予期せぬエラーが発生しました",
		statusCode: http.StatusInternalServerError,
	},
}
