package registry

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/usecase"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/infra/repository"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/interface/gen/openapi"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/interface/handler"
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
	return openapi.Handler(h), nil
}
