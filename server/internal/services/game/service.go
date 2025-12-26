package game

import (
	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) GetQuestionsByQuizID(quizID string) ([]models.Question, error) {
	var questions []models.Question
	if err := s.db.Where("quiz_id = ?", quizID).Preload("Options").Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}
