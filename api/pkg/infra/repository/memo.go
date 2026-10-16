package repository

import (
	"context"
	"database/sql"
	"errors"
	"uuid"

	appErr "github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/error"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/domain/entity"
)

type MemoRepository struct {
	db *sql.DB
}

func NewMemoRepository(db *sql.DB) *MemoRepository {
	return &MemoRepository{db: db}
}

func (r *MemoRepository) Insert(ctx context.Context, params *entity.Memo) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO memos (id, title, content, owner) VALUES ($1, $2, $3, $4)", params.ID, params.Title, params.Content, params.Owner)
	return err
}

func (r *MemoRepository) Update(ctx context.Context, params *entity.Memo) error {
	_, err := r.db.ExecContext(ctx, "UPDATE memos SET title = $1, content = $2, updated_at = NOW() WHERE id = $3 AND owner = $4", params.Title, params.Content, params.ID, params.Owner)
	return err
}

func (r *MemoRepository) Delete(ctx context.Context, id uuid.UUID, owner string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM memos WHERE id = $1 AND owner = $2", id, owner)
	return err
}

func (r *MemoRepository) FindByID(ctx context.Context, id uuid.UUID, owner string) (*entity.Memo, error) {
	var m entity.Memo
	res := r.db.QueryRowContext(ctx, "SELECT id, title, content, owner FROM memos WHERE id = $1 AND owner = $2", id, owner)
	if res.Err() != nil {
		return nil, res.Err()
	}
	err := res.Scan(&m.ID, &m.Title, &m.Content, &m.Owner)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, appErr.ErrNotFound
		}
		return nil, err
	}
	return &m, nil
}
