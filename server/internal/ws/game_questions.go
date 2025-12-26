package ws

import (
	"net/http"

	"github.com/datmedevil17/kahoot-quiz-go/internal/database"
	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
	"github.com/datmedevil17/kahoot-quiz-go/internal/utils"
	"github.com/gin-gonic/gin"
)


func GetQuizQuestions(c *gin.Context) {
	quizID := c.Param("quizId")

	// Optional: validate quiz exists
	var quiz models.Quiz
	if err := database.DB.First(&quiz, "id = ?", quizID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Quiz not found")
		return
	}

	var questions []models.Question
	if err := database.DB.
		Where("quiz_id = ?", quizID).
		Preload("Options").
		Order("created_at ASC").
		Find(&questions).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch questions")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Questions fetched successfully", gin.H{"questions": questions})
}
