package models

type PlayerScore struct {
	BaseModel

	GameSessionID string `gorm:"index"`
	PlayerID      string `gorm:"index"`
	Score         int
	Correct       int
	Incorrect     int
}
