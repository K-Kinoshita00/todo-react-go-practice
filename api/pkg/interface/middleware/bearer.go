package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/authservice"
	appErr "github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/error"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/interface/presenter"
)

type ctxKey struct{}

var claimsKey ctxKey

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
			claims, err := auth.Verify(r.Context(), token)
			if err != nil {
				appError := presenter.MapToAppError(r.Context(), err)
				presenter.NewResponse(appError.StatusCode, appError).Send(w)
				return
			}
			// claimsをcontextに追加
			ctx := context.WithValue(r.Context(), claimsKey, claims)
			// その Request を context を使って加工したもので次のハンドラに渡す
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ClaimsFromContext(ctx context.Context) (*authservice.Claims, bool) {
	v := ctx.Value(claimsKey)
	claims, ok := v.(*authservice.Claims)
	return claims, ok
}

func ContextWithClaims(ctx context.Context, claims *authservice.Claims) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
}