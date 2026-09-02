package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/usecase"
)

type TodoCommandRepository struct {
	db *sql.DB
}

func NewTodoCommandRepository(db *sql.DB) *TodoCommandRepository {
	return &TodoCommandRepository{db: db}
}

func (r *TodoCommandRepository) Insert(ctx context.Context, params usecase.InsertTodo) (uuid.UUID, error) {
	var id uuid.UUID
	err := r.db.QueryRowContext(ctx, `INSERT INTO todos (title, status) VALUES ($1, $2) RETURNING id`, params.Title, params.Status).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (r *TodoCommandRepository) Update(ctx context.Context, params usecase.UpdateTodo) error {
	_, err := r.db.ExecContext(ctx, `UPDATE todos SET title = $1, status = $2, updated_at = NOW() WHERE id = $3`, params.Title, params.Status, params.ID)
	return err
}

func (r *TodoCommandRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM todos WHERE id = $1`, id)
	return err
}
