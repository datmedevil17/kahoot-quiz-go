package quiz

import (
	"errors"

	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateQuiz(title, description string, createdBy uint) (*models.Quiz, error) {
	quiz := models.Quiz{
		Title:       title,
		Description: description,
		CreatedBy:   createdBy,
	}

	if err := s.db.Create(&quiz).Error; err != nil {
		return nil, err
	}

	return &quiz, nil
}

func (s *Service) GetQuizzesByUserID(userID uint) ([]models.Quiz, error) {
	var quizzes []models.Quiz
	if err := s.db.Where("created_by = ?", userID).Find(&quizzes).Error; err != nil {
		return nil, err
	}
	return quizzes, nil
}

func (s *Service) GetQuizByID(id string) (*models.Quiz, error) {
	var quiz models.Quiz
	if err := s.db.Preload("Questions.Options").First(&quiz, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (s *Service) ValidateOwnership(quizID string, userID uint) error {
	var quiz models.Quiz
	if err := s.db.Select("created_by").First(&quiz, "id = ?", quizID).Error; err != nil {
		return errors.New("quiz not found")
	}
	if quiz.CreatedBy != userID {
		return errors.New("not quiz owner")
	}
	return nil
}
