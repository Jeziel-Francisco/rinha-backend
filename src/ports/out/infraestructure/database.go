package infraestructure

import "context"

type Database interface {
	Connect(ctx context.Context, connectionString string) error
	Execution(query string, args ...interface{}) error
	Query(query string, args ...interface{}) (interface{}, error)
	Close(ctx context.Context) error
}
