package models

type Quiz struct {
	BaseModel

	Title       string     `gorm:"not null"`
	Description string
	CreatedBy   uint       `gorm:"not null"` // User ID
	Questions   []Question `gorm:"foreignKey:QuizID"`
}
