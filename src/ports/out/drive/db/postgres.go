package db

import (
	"context"
	"docker-example/src/commons"
	"docker-example/src/ports/out/infraestructure"
	"fmt"
	"os"
)

func NewPostgres() infraestructure.Database {
	postgres := infraestructure.NewPostgresSql()

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_POSTGRES_USER"),
		os.Getenv("DB_POSTGRES_PASSWORD"),
		os.Getenv("DB_POSTGRES_HOST"),
		os.Getenv("DB_POSTGRES_PORT"),
		os.Getenv("DB_POSTGRES_DB"),
	)

	if err := postgres.Connect(context.Background(), connectionString); err != nil {
		panic(err)
	}
	defer postgres.Close(context.Background())

	commons.ContainerInjectable.Register(commons.PostgresDatabase, postgres)

	return postgres
}
