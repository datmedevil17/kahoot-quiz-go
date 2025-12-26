package ws

import "time"

func startTimer(room *Room, seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)

	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	if !room.Started {
		return
	}

	if room.CurrentQ >= len(room.Questions) {
		return
	}

	q := room.Questions[room.CurrentQ]

	for clientID, option := range room.Answers {
		isCorrect := option == q.Answer

		elapsed := time.Since(room.StartTime)
		score := calculateScore(isCorrect, elapsed, q.TimeLimit)

		room.Scores[clientID] += score
	}

	broadcastLeaderboard(room)

	room.CurrentQ++

	time.Sleep(2 * time.Second)
	sendNextQuestion(room)
}
