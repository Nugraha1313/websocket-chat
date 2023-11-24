package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Nugraha1313/websocket-chat/internal/handler"
)

func NewHttpServer() *http.Server {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/ws", handler.WebSocketConnect)

	return &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
}
