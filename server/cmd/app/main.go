package main

import "github.com/Nugraha1313/websocket-chat/internal/server"

func main() {
	s := server.NewHttpServer()
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
