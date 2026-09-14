package authservice

import (
	"context"
)

type Claims struct {
	Sub string
}

type AuthService interface {
	// Verify はトークンを検証し、claims を返す
	Verify(ctx context.Context, token string) (*Claims, error)
}
