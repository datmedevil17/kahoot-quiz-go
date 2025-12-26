package ws

import (
	"encoding/json"

	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
)

func readPump(client *Client, hub *Hub) {
	defer func() {
		if client.Room != nil {
			handleDisconnect(client, hub)
		}
		client.Conn.Close()
	}()

	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			break
		}

		var msg models.WSMessage
		if err := json.Unmarshal(message, &msg); err != nil {
			continue
		}

		handleEvent(client, hub, msg)
	}
}
