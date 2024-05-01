package infraestructure

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database interface {
	Connect(ctx context.Context, connectionString string) error
	Close(ctx context.Context) error
	GetPoolConnection() *pgxpool.Pool
	HasError(err error) bool
	HasEmptyData(err error) bool
}
