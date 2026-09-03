package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/domain/entity"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Insert(ctx context.Context, params *entity.Todo) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO todos (id,title, status) VALUES ($1, $2, $3)`, params.ID, params.Title, params.Status)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) Update(ctx context.Context, params *entity.Todo) error {
	_, err := r.db.ExecContext(ctx, `UPDATE todos SET title = $1, status = $2, updated_at = NOW() WHERE id = $3`, params.Title, params.Status, params.ID)
	return err
}

func (r *TodoRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM todos WHERE id = $1`, id)
	return err
}

func (r *TodoRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Todo, error) {
	var t entity.Todo
	res := r.db.QueryRowContext(ctx, `SELECT id, title, status FROM todos WHERE id = $1`, id)
	if res.Err() != nil {
		return nil, res.Err()
	}
	err := res.Scan(&t.ID, &t.Title, &t.Status)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
