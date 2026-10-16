package dto

import (
	"uuid"
)

type Memo struct {
	ID      uuid.UUID
	Title   string
	Content string
}
