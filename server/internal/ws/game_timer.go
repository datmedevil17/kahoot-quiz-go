package ws

import "time"

func startTimer(room *Room, seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)

	room.Mutex.Lock()
	broadcastLeaderboard(room)

	room.CurrentQ++
	room.Mutex.Unlock()

	time.Sleep(2 * time.Second)
	sendNextQuestion(room)
}
