package models

type Option struct {
	BaseModel

	QuestionID string `gorm:"index;not null"`
	Text       string `gorm:"not null"`
}
