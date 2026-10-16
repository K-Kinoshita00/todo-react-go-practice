package entity

import (
	"errors"
	"uuid"
)

type Memo struct {
	ID      uuid.UUID
	Title   string
	Content string
	Owner   string
}

func NewMemo(
	title string,
	content string,
	owner string,
) (*Memo, error) {
	return &Memo{
		ID:      uuid.New(),
		Title:   title,
		Content: content,
		Owner:   owner,
	}, nil
}

func (e *Memo) UpdateTitleAndContent(
	title string,
	content string,
) error {
	e.Title = title
	e.Content = content
	return nil
}

func (e *Memo) Validate() error {
	if e.Title == "" {
		return errors.New("title is required")
	}
	if e.Owner == "" {
		return errors.New("owner is required")
	}
	return nil
}
