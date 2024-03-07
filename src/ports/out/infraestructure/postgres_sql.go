package infraestructure

import (
	"context"
)

func NewPostgresSql() Database {
	return &postgresSql{}
}

type postgresSql struct {
}

func (p *postgresSql) Connect(ctx context.Context, connectionString string) error {
	return nil
}

func (p *postgresSql) Execution(query string, args ...interface{}) error {
	return nil
}

func (p *postgresSql) Query(query string, args ...interface{}) (interface{}, error) {
	return nil, nil
}

func (p *postgresSql) Close(ctx context.Context) error {
	return nil
}
