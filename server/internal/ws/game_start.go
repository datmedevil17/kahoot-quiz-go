package ws

import (
	"encoding/json"

	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
)

func handleStartGame(client *Client, hub *Hub, data interface{}) {
	room := client.Room
	if room == nil || room.Started {
		return
	}

	// Only host can start
	if client.UserID != room.HostID {
		sendError(client, "Only host can start the game")
		return
	}

	payloadBytes, _ := json.Marshal(data)
	var payload models.StartGamePayload
	json.Unmarshal(payloadBytes, &payload)

	questions, err := hub.GameService.GetQuestionsByQuizID(payload.QuizID)
	if err != nil {
		sendError(client, "Failed to load questions")
		return
	}

	room.Questions = questions
	room.CurrentQ = 0
	room.Started = true
	room.Scores = make(map[string]int)
	room.QuizID = payload.QuizID

	sendNextQuestion(room)
}
