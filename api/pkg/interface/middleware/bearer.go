package middleware

import (
	"net/http"
	"strings"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/authservice"
	appErr "github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/error"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/interface/presenter"
)

func Bearer(auth authservice.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// ヘルスチェックの場合はスキップ
			if r.URL.Path == "/health" {
				next.ServeHTTP(w, r)
				return
			}
			// ヘッダーのAuthorizationを取得
			raw := r.Header.Get("Authorization")
			// "Bearer"を外してトークンを取得
			token, ok := strings.CutPrefix(raw, "Bearer ")
			if !ok || token == "" {
				appError := presenter.MapToAppError(r.Context(), appErr.ErrUnauthorized)
				presenter.NewResponse(appError.StatusCode, appError).Send(w)
				return
			}
			_, err := auth.Verify(r.Context(), token)
			if err != nil {
				appError := presenter.MapToAppError(r.Context(), err)
				presenter.NewResponse(appError.StatusCode, appError).Send(w)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}