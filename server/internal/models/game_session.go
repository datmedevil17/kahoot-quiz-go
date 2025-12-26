package models

type GameSession struct {
	BaseModel

	QuizID   string `gorm:"index"`
	HostID   string `gorm:"index"`
	GamePIN  string `gorm:"uniqueIndex"`
	Ended    bool
}
