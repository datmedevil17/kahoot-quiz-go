package ws

import (
	"sort"

	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
)

func broadcastLeaderboard(room *Room) {
	var leaderboard []models.LeaderboardEntry

	for _, client := range room.Clients {
		// Skip Host in leaderboard
		if client.UserID == room.HostID {
			continue
		}

		leaderboard = append(leaderboard, models.LeaderboardEntry{
			Username: client.Username,
			Score:    room.Scores[client.ID],
		})
	}

	sort.Slice(leaderboard, func(i, j int) bool {
		return leaderboard[i].Score > leaderboard[j].Score
	})

	broadcast(room, "leaderboard", leaderboard)
}
