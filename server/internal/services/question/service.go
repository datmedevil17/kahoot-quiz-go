package question

import (
	"strconv"

	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateQuestion(quizID string, text string, answer int, timeLimit int, options []string) (*models.Question, []models.Option, error) {
	question := models.Question{
		QuizID:    quizID,
		Text:      text,
		Answer:    answer,
		TimeLimit: timeLimit,
	}

	if err := s.db.Create(&question).Error; err != nil {
		return nil, nil, err
	}

	var optionModels []models.Option
	for _, opt := range options {
		optionModels = append(optionModels, models.Option{
			QuestionID: strconv.FormatUint(uint64(question.ID), 10),
			Text:       opt,
		})
	}

	if err := s.db.Create(&optionModels).Error; err != nil {
		return nil, nil, err
	}

	return &question, optionModels, nil
}

func (s *Service) GetQuestionsByQuizID(quizID string) ([]models.Question, error) {
	var questions []models.Question
	if err := s.db.Where("quiz_id = ?", quizID).Preload("Options").Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}
