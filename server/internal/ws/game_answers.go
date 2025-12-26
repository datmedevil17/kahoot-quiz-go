package ws

import (
	"encoding/json"

	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
)

func handleSubmitAnswer(client *Client, data interface{}) {
	room := client.Room
	if room == nil || !room.Started {
		return
	}

	payloadBytes, _ := json.Marshal(data)
	var payload models.AnswerSubmissionPayload
	json.Unmarshal(payloadBytes, &payload)

	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	// One answer per question
	if _, exists := room.Answers[client.ID]; exists {
		return
	}

	room.Answers[client.ID] = payload.Option
}
