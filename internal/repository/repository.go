package repository

import (
	"context"
	"database/sql"
)

type Repository[T any] interface {
	Create(ctx context.Context, tx *sql.Tx, entity *T) error
	Update(ctx context.Context, tx *sql.Tx, entity *T) error
	Delete(ctx context.Context, tx *sql.Tx, entity *T) error
	FindById(ctx context.Context, tx *sql.Tx, entity *T) error
	FindAll(ctx context.Context, tx *sql.Tx) ([]T, error)
}
