package repository

import (
	"context"
	"database/sql"
	"errors"
	"uuid"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/dto"
	appErr "github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/error"
)

type MemoQueryRepository struct {
	db *sql.DB
}

func NewMemoQueryRepository(db *sql.DB) *MemoQueryRepository {
	return &MemoQueryRepository{db: db}
}

func (r *MemoQueryRepository) List(ctx context.Context, owner string) ([]*dto.Memo, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, title, content FROM memos WHERE owner = $1", owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var memos []*dto.Memo
	for rows.Next() {
		var m dto.Memo
		err := rows.Scan(&m.ID, &m.Title, &m.Content)
		if err != nil {
			return nil, err
		}
		memos = append(memos, &m)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return memos, nil
}

func (r *MemoQueryRepository) FindByID(ctx context.Context, id uuid.UUID, owner string) (*dto.Memo, error) {
	var m dto.Memo
	res := r.db.QueryRowContext(ctx, "SELECT id, title, content FROM memos WHERE id = $1 AND owner = $2", id, owner)
	if res.Err() != nil {
		return nil, res.Err()
	}
	err := res.Scan(&m.ID, &m.Title, &m.Content)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, appErr.ErrNotFound
		}
		return nil, err
	}
	return &m, nil
}
