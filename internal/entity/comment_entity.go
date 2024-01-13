package entity

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	Id        uuid.UUID
	Body      uuid.UUID
	CreatedAt time.Time
}
