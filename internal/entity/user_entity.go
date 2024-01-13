package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID
	Username  string
	Password  string
	Email     string
	Phone     string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
