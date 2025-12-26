package utils

import (
	"errors"

	"github.com/datmedevil17/kahoot-quiz-go/internal/database"
	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
)

func ValidateQuizOwnership(quizID string, userID uint) (*models.Quiz, error) {
	var quiz models.Quiz
	if err := database.DB.First(&quiz, "id = ?", quizID).Error; err != nil {
		return nil, errors.New("quiz not found")
	}

	if quiz.CreatedBy != userID {
		return nil, errors.New("not quiz owner")
	}

	return &quiz, nil
}
