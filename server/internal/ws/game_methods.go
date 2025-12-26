package ws

import (
	"fmt"
	"time"

	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
)

func sendNextQuestion(room *Room) {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	// Check if we've run out of questions
	if room.CurrentQ >= len(room.Questions) {
		// Broadcast game over (IO inside lock, acceptable for this scale to ensure ordering)
		broadcast(room, "GAME_OVER", room.Scores)
		room.Started = false

		// Persistence should be outside lock if possible, or inside if it's fast.
		// persistGameResults db call might be slow.
		// Spawning goroutine is fine, but we need to ensure room state doesn't change strangely in between.
		// Passed room pointer is fine.
		go func() {
			persistGameResults(room)
			if room.Hub != nil {
				room.Hub.RemoveRoom(room.PIN)
			}
		}()
		return
	}

	q := room.Questions[room.CurrentQ]

	// Reset answers for the new question
	room.Answers = make(map[string]int)

	options := make([]string, len(q.Options))
	for i, opt := range q.Options {
		options[i] = opt.Text
	}

	payload := models.QuestionPayload{
		ID:        fmt.Sprint(q.ID),
		Text:      q.Text,
		Options:   options,
		TimeLimit: q.TimeLimit,
	}

	broadcast(room, "NEXT_QUESTION", payload)

	room.StartTime = time.Now()

	// Run timer in a separate goroutine to avoid recursion/blocking
	go startTimer(room, q.TimeLimit)
}
