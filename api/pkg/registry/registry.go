package registry

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/authservice"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/usecase"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/infra/auth"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/infra/repository"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/interface/gen/openapi"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/interface/handler"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/interface/middleware"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/interface/presenter"
)

type Handler struct {
	*handler.HealthHandler
	*handler.TodoHandler
}

func NewHandler() *Handler {
	return &Handler{}
}

func getDSN() string {
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresDB := os.Getenv("POSTGRES_DB")
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", postgresUser, postgresPassword, postgresHost, postgresPort, postgresDB)
}

func NewRegistry() (http.Handler, error) {
	db, err := sql.Open("postgres", getDSN())
	if err != nil {
		return nil, err
	}

	cmd := repository.NewTodoRepository(db)
	query := repository.NewTodoQueryRepository(db)
	uc := usecase.NewTodoUseCase(cmd, query)
	h := &Handler{
		HealthHandler: handler.NewHealthHandler(),
		TodoHandler:   handler.NewTodoHandler(uc),
	}
	var authMode = os.Getenv("AUTH_MODE")
	var a authservice.AuthService
	if authMode == "hs256" {
		// JWT_SECRETを使ってHS256を初期化
		a = auth.NewHS256(os.Getenv("JWT_SECRET"))
	} else {
		// JWKSを初期化
		a = auth.NewJWKS(
			os.Getenv("JWKS_URL"),
			os.Getenv("JWT_ISSUER"),
			os.Getenv("JWT_AUDIENCE"),
		)
	}
	options := openapi.StdHTTPServerOptions{
		BaseURL:          "",
		BaseRouter:       nil,
		Middlewares:      nil,
		ErrorHandlerFunc: presenter.BindError,
	}
	bearerHandler := middleware.Bearer(a)(openapi.HandlerWithOptions(h, options))
	corsHandler := middleware.CORS(os.Getenv("CORS_ORIGIN"))(bearerHandler)
	securityHeadersHandler := middleware.SecurityHeaders()(corsHandler)
	return securityHeadersHandler, nil
}
