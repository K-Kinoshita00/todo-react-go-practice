package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/queryservice"
)

type TodoQueryRepository struct {
	db *sql.DB
}

func NewTodoQueryRepository(db *sql.DB) *TodoQueryRepository {
	return &TodoQueryRepository{db: db}
}

func (r *TodoQueryRepository) List(ctx context.Context) ([]*queryservice.Todo, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, title, status, created_at, updated_at FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*queryservice.Todo
	for rows.Next() {
		var t queryservice.Todo
		err := rows.Scan(&t.ID, &t.Title, &t.Status, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &t)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return todos, nil
}

func (r *TodoQueryRepository) FindByID(ctx context.Context, id uuid.UUID) (*queryservice.Todo, error) {
	var t queryservice.Todo
	res := r.db.QueryRowContext(ctx, `SELECT id, title, status, created_at, updated_at FROM todos WHERE id = $1`, id)
	if res.Err() != nil {
		return nil, res.Err()
	}
	err := res.Scan(&t.ID, &t.Title, &t.Status, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}