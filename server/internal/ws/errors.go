package ws

import (
	"encoding/json"

	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
	"github.com/gin-gonic/gin"
)

func sendError(client *Client, message string) {
	msg, _ := json.Marshal(models.WSMessage{
		Event: "error",
		Data: gin.H{
			"message": message,
		},
	})

	client.Send <- msg
}
