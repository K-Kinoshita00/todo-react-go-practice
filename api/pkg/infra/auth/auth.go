package auth

import (
	"context"
	"errors"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/authservice"
	appErr "github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/error"
)

type HS256 struct {
	secret []byte
}

func NewHS256(secret string) *HS256 {
	return &HS256{secret: []byte(secret)}
}

func (a *HS256) Verify(ctx context.Context, token string) (*authservice.Claims, error) {
	claims := jwt.MapClaims{}
	// ParseWithClaimsはトークンを解析しトークンを返す、鍵はa.secret
	parsed, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, jwt.ErrInvalidKeyType
		}
		return a.secret, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, appErr.ErrBearerTokenExpired
		}
		return nil, appErr.ErrUnauthorized
	}
	if !parsed.Valid {
		return nil, appErr.ErrUnauthorized
	}
	if claims["sub"] == nil {
		return nil, appErr.ErrUnauthorized
	}
	return &authservice.Claims{
		Sub: claims["sub"].(string),
	}, nil
}
