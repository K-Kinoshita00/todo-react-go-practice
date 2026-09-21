package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/authservice"
	appErr "github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/error"
)

type fakeAuthService struct {
	claims *authservice.Claims
	err    error
}

func (f *fakeAuthService) Verify(ctx context.Context, token string) (*authservice.Claims, error) {
	return f.claims, f.err
}

func TestBearer_OK(t *testing.T) {
	authSub := "test_user"
	auth := &fakeAuthService{
		claims: &authservice.Claims{
			Sub: authSub,
		},
		err: nil,
	}
	var (
		called bool
		got    *authservice.Claims
		ok     bool
	)
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		got, ok = ClaimsFromContext(r.Context())
	})
	bearerHandler := Bearer(auth)(next)
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	req.Header.Set("Authorization", "Bearer dummy")
	bearerHandler.ServeHTTP(httptest.NewRecorder(), req)

	// NOTE: bearer が通った後に next が呼ばれる
	if !called {
		t.Fatalf("next is not called")
	}
	if !ok {
		t.Fatalf("ClaimsFromContext is not ok")
	}
	if got.Sub != authSub {
		t.Fatalf("Claims Sub is not test_user")
	}
}

func TestBearer_OK_Health(t *testing.T) {
	auth := &fakeAuthService{}
	var called bool
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	})
	bearerHandler := Bearer(auth)(next)
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	bearerHandler.ServeHTTP(rec, req)

	if !called {
		t.Fatalf("next is not called")
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("recorder is not ok: got %d want %d", rec.Code, http.StatusOK)
	}
}

// ヘッダーなし
func TestBearer_Unauthorized_EmptyHeader(t *testing.T) {
	auth := &fakeAuthService{}
	var called bool
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	})
	bearerHandler := Bearer(auth)(next)

	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec := httptest.NewRecorder()
	bearerHandler.ServeHTTP(rec, req)

	if called {
		t.Fatalf("next is called")
	}
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("Error is not ErrUnauthorized: got %d want %d", rec.Code, http.StatusUnauthorized)
	}
}

// token が異なる場合
func TestBearer_Unauthorized_InvalidToken(t *testing.T) {
	auth := &fakeAuthService{
		err: appErr.ErrUnauthorized,
	}
	var called bool
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	})
	bearerHandler := Bearer(auth)(next)

	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec := httptest.NewRecorder()
	req.Header.Set("Authorization", "Bearer dummy")
	bearerHandler.ServeHTTP(rec, req)
	if called {
		t.Fatalf("next is called")
	}
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("Error is not ErrUnauthorized: got %d want %d", rec.Code, http.StatusUnauthorized)
	}
}
