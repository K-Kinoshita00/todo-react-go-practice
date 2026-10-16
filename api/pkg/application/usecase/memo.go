package usecase

import (
	"context"
	"uuid"

	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/dto"
	appErr "github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/error"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/application/queryservice"
	"github.com/K-Kinoshita00/todo-react-go-practice/pkg/domain/entity"
)

type MemoUseCase struct {
	cmd   MemoRepository
	query queryservice.MemoQueryService
}

func NewMemoUseCase(cmd MemoRepository, query queryservice.MemoQueryService) *MemoUseCase {
	return &MemoUseCase{cmd, query}
}

func (u *MemoUseCase) Create(ctx context.Context, title string, content string, owner string) error {
	ent, err := entity.NewMemo(title, content, owner)
	if err != nil || ent == nil {
		return err
	}
	if res := ent.Validate(); res != nil {
		return appErr.ErrBadRequest
	}
	if err = u.cmd.Insert(ctx, ent); err != nil {
		return err
	}
	return nil
}

func (u *MemoUseCase) Update(ctx context.Context, id uuid.UUID, title string, content string, owner string) error {
	ent, err := u.cmd.FindByID(ctx, id, owner)
	if err != nil {
		return err
	}
	if ent == nil {
		return appErr.ErrNotFound
	}
	if err = ent.UpdateTitleAndContent(title, content); err != nil {
		return appErr.ErrBadRequest
	}
	if res := ent.Validate(); res != nil {
		return appErr.ErrBadRequest
	}
	if err = u.cmd.Update(ctx, ent); err != nil {
		return err
	}
	return nil
}

func (u *MemoUseCase) Delete(ctx context.Context, id uuid.UUID, owner string) error {
	err := u.cmd.Delete(ctx, id, owner)
	if err != nil {
		return err
	}
	return nil
}

func (u *MemoUseCase) List(ctx context.Context, owner string) ([]*dto.Memo, error) {
	memos, err := u.query.List(ctx, owner)
	if err != nil {
		return nil, err
	}
	return memos, err
}
func (u *MemoUseCase) FindByID(ctx context.Context, id uuid.UUID, owner string) (*dto.Memo, error) {
	memo, err := u.query.FindByID(ctx, id, owner)
	if err != nil {
		return nil, err
	}
	if memo == nil {
		return nil, appErr.ErrNotFound
	}
	return memo, nil
}

type MemoRepository interface {
	Insert(ctx context.Context, params *entity.Memo) error
	Update(ctx context.Context, params *entity.Memo) error
	Delete(ctx context.Context, id uuid.UUID, owner string) error
	FindByID(ctx context.Context, id uuid.UUID, owner string) (*entity.Memo, error)
}
