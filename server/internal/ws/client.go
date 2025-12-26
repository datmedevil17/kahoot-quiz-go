package ws

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	ID       string
	UserID   string
	Email    string
	Username string
	Conn     *websocket.Conn
	Room     *Room
	Send     chan []byte
}
