package models

type LeaderboardEntry struct {
	Username string `json:"username"`
	Score    int    `json:"score"`
}
