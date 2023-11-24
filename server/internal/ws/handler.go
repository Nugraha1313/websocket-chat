package ws

import (
	"github.com/Nugraha1313/websocket-chat/internal/array"
	"github.com/Nugraha1313/websocket-chat/internal/client"
	"github.com/Nugraha1313/websocket-chat/internal/model"
)

func MemberJoin(clients client.WebSocketClientsPool, cl client.WebSocketClient) {
	broadcast(clients, cl, model.WebSocketMessage{
		Type: MEMBER_JOIN,
		Content: map[string]string{
			"id": cl.Id(),
		},
	})
}

func MemberLeave(clients client.WebSocketClientsPool, cl client.WebSocketClient) {
	broadcast(clients, cl, model.WebSocketMessage{
		Type: MEMBER_LEAVE,
		Content: map[string]string{
			"id": cl.Id(),
		},
	})
}

func NewMessage(clients client.WebSocketClientsPool, cl client.WebSocketClient, message string) {
	broadcast(clients, cl, model.WebSocketMessage{
		Type: MESSAGE,
		Content: map[string]string{
			"id":      cl.Id(),
			"message": message,
		},
	})
}

func broadcast(clients client.WebSocketClientsPool, cl client.WebSocketClient, msg model.WebSocketMessage) {
	array.ForEach(
		array.Except(clients, func(item client.WebSocketClient) bool { return item.Id() == cl.Id() }),
		func(item client.WebSocketClient) { item.Write(msg) },
	)
}
