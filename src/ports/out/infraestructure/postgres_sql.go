package infraestructure

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresSql() Database {
	return &postgresSql{}
}

type postgresSql struct {
	poolConnection *pgxpool.Pool
}

func (p *postgresSql) Connect(ctx context.Context, connectionString string) error {
	poolConnection, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		fmt.Printf("Opened error connection: %s\n", err.Error())
		return err
	}
	fmt.Printf("Opened connection")
	p.poolConnection = poolConnection
	return nil
}

func (p *postgresSql) GetPoolConnection() *pgxpool.Pool {
	return p.poolConnection
}

func (p *postgresSql) HasError(err error) bool {
	return err != nil && err != pgx.ErrNoRows
}

func (p *postgresSql) HasEmptyData(err error) bool {
	return err != nil && err == pgx.ErrNoRows
}

func (p *postgresSql) Execution(query string, args ...interface{}) (interface{}, error) {
	result, err := p.poolConnection.Exec(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *postgresSql) QueryRow(query string, args []any, outputArg ...any) error {
	err := p.poolConnection.QueryRow(context.Background(), query, args...).Scan(outputArg...)
	// p.poolConnection.QueryRow(context.Background(), query, "jeize").Scan(outputArg)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgresSql) Close(ctx context.Context) error {
	p.poolConnection.Close()
	return nil
}
