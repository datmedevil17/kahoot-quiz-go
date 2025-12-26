package ws

import (
	"encoding/json"

	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
)

func broadcast(room *Room, event string, data interface{}) {
	msg, _ := json.Marshal(models.WSMessage{
		Event: event,
		Data:  data,
	})

	for _, client := range room.Clients {
		client.Send <- msg
	}
}
