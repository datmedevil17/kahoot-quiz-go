package ws

import "time"

const (
	BasePoints  = 1000
	MaxBonus    = 500
)

func calculateScore(
	isCorrect bool,
	elapsed time.Duration,
	timeLimit int,
) int {
	if !isCorrect {
		return 0
	}

	remaining := float64(timeLimit) - elapsed.Seconds()
	if remaining < 0 {
		remaining = 0
	}

	bonus := int((remaining / float64(timeLimit)) * MaxBonus)
	return BasePoints + bonus
}
