package middleware

import "net/http"

func SecurityHeaders() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Content-Type-Options", "nosniff") // ブラウザに MIME を推測させない
			w.Header().Set("X-Frame-Options", "DENY") // iframe に載せない
			next.ServeHTTP(w, r)
		})
	}
}