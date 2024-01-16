package repository

import (
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, tx *sql.Tx, entity any) error
	Update(ctx context.Context, tx *sql.Tx, entity any) error
	Delete(ctx context.Context, tx *sql.Tx, entity any) error
	FindById(ctx context.Context, tx *sql.Tx, entity any) error
	FindAll(ctx context.Context, tx *sql.Tx) ([]any, error)
}
