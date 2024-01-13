package entity

import (
	"github.com/google/uuid"
	"time"
)

type Article struct {
	Id          uuid.UUID
	Thumbnail   string
	Title       string
	Description string
	Body        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserId      uuid.UUID
}
