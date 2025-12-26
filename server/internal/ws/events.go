package ws

import (
	"encoding/json"

	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
	"github.com/gin-gonic/gin"
)

func handleEvent(client *Client, hub *Hub, msg models.WSMessage) {
	switch msg.Event {
	case "join_game":
		handleJoinGame(client, hub, msg.Data)
	}
}

func handleJoinGame(client *Client, hub *Hub, data interface{}) {
	payloadBytes, _ := json.Marshal(data)

	var payload models.JoinGamePayload
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return
	}

	room, ok := hub.Rooms[payload.GamePIN]
	if !ok {
		sendError(client, "Invalid game PIN")
		return
	}

	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	// Username now comes from auth (email or derived)
	// We trust client.Username which was set from JWT
	// Actually, let's ensure it's set from auth if not already
	username := client.Username

	// Check for reconnect
	for _, c := range room.Clients {
		if c.UserID == client.UserID {
			// Replace old connection
			c.Conn.Close()
			delete(room.Clients, c.ID)
		}
	}

	// Prevent duplicate usernames from DIFFERENT users
	for _, c := range room.Clients {
		if c.Username == username {
			sendError(client, "User already joined")
			return
		}
	}

	if room.Started {
		sendError(client, "Game already started")
		return
	}

	client.Room = room
	room.Clients[client.ID] = client

	broadcast(room, "player_joined", gin.H{
		"username": username,
	})
}
