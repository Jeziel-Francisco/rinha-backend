package main

import (
	"docker-example/src/ports/in/drive/web"
	"docker-example/src/ports/in/handler"
	"docker-example/src/ports/out/drive/db"
)

func main() {
	injectableDependencies()
	initializeGinApi()
}

func injectableDependencies() {
	// startup database
	db.NewPostgres()

	// startup repository database
}

func initializeGinApi() {
	handler := handler.NewHandler()
	web.NewGin(handler).InitializeGin()
}
