package models

type Question struct {
	BaseModel

	QuizID   string   `gorm:"index;not null"`
	Text     string   `gorm:"not null"`
	Options  []Option `gorm:"foreignKey:QuestionID"`
	Answer   int      `gorm:"not null"` // index of correct option
	TimeLimit int     `gorm:"not null"` // seconds
}
