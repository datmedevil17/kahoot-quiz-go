package ws

import (
	"log"
	"strconv"

	"github.com/datmedevil17/kahoot-quiz-go/internal/database"
	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
)

// Persist Game Session
func persistGameResults(room *Room) {
	// Create game session record
	session := models.GameSession{
		QuizID:  room.QuizID,
		HostID:  room.HostID,
		GamePIN: room.PIN,
		Ended:   true,
	}

	if err := database.DB.Create(&session).Error; err != nil {
		log.Println("Failed to save game session:", err)
		return
	}

	// Save player scores
	for clientID, score := range room.Scores {
		ps := models.PlayerScore{
			GameSessionID: strconv.FormatUint(uint64(session.ID), 10),
			PlayerID:      clientID,
			Score:         score,
		}

		if err := database.DB.Create(&ps).Error; err != nil {
			log.Println("Failed to save player score:", err)
		}
	}
}
