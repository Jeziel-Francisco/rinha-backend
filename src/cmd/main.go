package main

import (
	"docker-example/src/application/domain/service"
	"docker-example/src/application/usecase"
	"docker-example/src/ports/in/drive/web"
	"docker-example/src/ports/in/handler"
	personDbAdapter "docker-example/src/ports/out/drive/adapter/person"
	db "docker-example/src/ports/out/drive/database"
	personDb "docker-example/src/ports/out/drive/database/person"
	"os"
)

func main() {
	if os.Getenv("ENVIRONMENT") == "LOCAL" || os.Getenv("ENVIRONMENT") == "" {
		os.Setenv("DB_POSTGRES_USER", "postgres")
		os.Setenv("DB_POSTGRES_PASSWORD", "postgres")
		os.Setenv("DB_POSTGRES_HOST", "localhost")
		os.Setenv("DB_POSTGRES_PORT", "5432")
		os.Setenv("DB_POSTGRES_DB", "docker_example")
	}
	handler := initializeDependencies()
	initializeGinApi(handler)
}

func initializeDependencies() handler.Handler {
	// database dependencies
	databaseConnection := db.NewPostgres()

	// drive output client dependencies
	personDatabase := personDb.NewPersonDatabase(databaseConnection)

	// drive output adapter dependencies
	personAdapter := personDbAdapter.NewPersonDatabaseAdapter(personDatabase)

	// application service
	checkExistsnickNameService := service.NewCheckExistsNicknameService(personAdapter)
	personCreateService := service.NewPersonCreateService(personAdapter)

	// application usecase
	personCreateUseCase := usecase.NewPersonCreateUseCase(checkExistsnickNameService, personCreateService)

	return handler.NewHandler(personCreateUseCase)

}

func initializeGinApi(handler handler.Handler) {
	web.NewGin(handler).InitializeGin()
}
