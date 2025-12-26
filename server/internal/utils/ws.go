package utils

type WSMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

type JoinGamePayload struct {
	GamePIN  string `json:"game_pin"`
	Username string `json:"username"`
}


type AnswerPayload struct {
	QuestionID string `json:"question_id"`
	Option     int    `json:"option"`
}

type LeaderboardEntry struct {
	PlayerID string `json:"player_id"`
	Score    int    `json:"score"`
}