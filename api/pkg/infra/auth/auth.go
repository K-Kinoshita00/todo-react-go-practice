package auth

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/authservice"
	appErr "github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/error"
)

type HS256 struct {
	secret []byte
}

type JWKS struct {
	jwksURL  string
	issuer   string
	audience string
}

func NewHS256(secret string) *HS256 {
	return &HS256{secret: []byte(secret)}
}

func NewJWKS(jwksURL, issuer, audience string) *JWKS {
	return &JWKS{jwksURL: jwksURL, issuer: issuer, audience: audience}
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

func (a *JWKS) Verify(ctx context.Context, token string) (*authservice.Claims, error) {
	claims := jwt.MapClaims{}
	parsed, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodRS256 {
			return nil, jwt.ErrInvalidKeyType
		}
		// kidはJWTのヘッダーのkidフィールドの値で、公開鍵のIDを指定する
		kid, ok := t.Header["kid"].(string)
		if !ok {
			return nil, jwt.ErrInvalidKey
		}
		// JWKSのURLにGETリクエストを送信して公開鍵を取得する
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.jwksURL, nil)
		if err != nil {
			return nil, err
		}
		// リクエストを送信して公開鍵を取得する
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to get JWKS: %d", resp.StatusCode)
		}
		var jwkSet struct {
			Keys []struct {
				Kid string `json:"kid"`
				N   string `json:"n"`
				E   string `json:"e"`
			} `json:"keys"`
		}
		err = json.NewDecoder(resp.Body).Decode(&jwkSet)
		if err != nil {
			return nil, err
		}
		for _, key := range jwkSet.Keys {
			if key.Kid == kid {
				nBytes, err := base64.RawURLEncoding.DecodeString(key.N)
				if err != nil {
					return nil, err
				}
				eBytes, err := base64.RawURLEncoding.DecodeString(key.E)
				if err != nil {
					return nil, err
				}
				pub := &rsa.PublicKey{
					N: new(big.Int).SetBytes(nBytes),
					E: int(new(big.Int).SetBytes(eBytes).Int64()),
				}
				return pub, nil
			}
		}
		return nil, jwt.ErrInvalidKey
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
	switch aud := claims["aud"].(type) {
	case string:
		if aud != a.audience {
			return nil, appErr.ErrUnauthorized
		}
	case []any:
		ok := false
		for _, v := range aud {
			s, isString := v.(string)
			if isString && s == a.audience {
				ok = true
				break
			}
		}
		if !ok {
			return nil, appErr.ErrUnauthorized
		}
	default:
		return nil, appErr.ErrUnauthorized
	}

	if claims["iss"] != a.issuer {
		return nil, appErr.ErrUnauthorized
	}
	return &authservice.Claims{
		Sub: claims["sub"].(string),
	}, nil
}
