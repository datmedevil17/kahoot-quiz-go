package ws

import "time"

func startTimer(room *Room, seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)

	room.Mutex.Lock()
	if room.Started && room.CurrentQ < len(room.Questions) {
		q := room.Questions[room.CurrentQ]

		for clientID, option := range room.Answers {
			isCorrect := option == q.Answer

			elapsed := time.Since(room.StartTime)
			score := calculateScore(isCorrect, elapsed, q.TimeLimit)

			room.Scores[clientID] += score
		}
	}
	
	broadcastLeaderboard(room)

	room.CurrentQ++
	room.Mutex.Unlock()

	time.Sleep(2 * time.Second)
	sendNextQuestion(room)
}
