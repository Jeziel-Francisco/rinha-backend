package web

import (
	"docker-example/src/ports/in/drive/adapter"
	"docker-example/src/ports/in/handler"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	handler handler.Handler
}

func NewGin(handler handler.Handler) *Gin {
	return &Gin{
		handler: handler,
	}
}

func (webGin Gin) InitializeGin() {
	r := gin.Default()
	r.GET("/ping", adapter.GinAdapter(webGin.handler.Ping))
	r.POST("/person", adapter.GinAdapter(webGin.handler.PersonCreate))
	r.Run()
}
