package quiz

import (
	"net/http"

	"github.com/datmedevil17/kahoot-quiz-go/internal/database"
	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
	"github.com/datmedevil17/kahoot-quiz-go/internal/utils"
	"github.com/gin-gonic/gin"
)

func CreateQuiz(c *gin.Context) {
	var req CreateQuizRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}
	userId, _ := c.Get("userID")

	quiz := models.Quiz{
		Title:       req.Title,
		Description: req.Description,
		CreatedBy:   uint(userId.(uint)),
	}

	if err := database.DB.Create(&quiz).Error; err != nil {
		utils.ErrorResponse(c, 500, "Failed to create quiz")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Quiz created successfully", gin.H{"quiz": quiz})
}

func GetMyQuizzes(c *gin.Context) {
	userId, _ := c.Get("userID")

	var quizzes []models.Quiz
	if err := database.DB.
		Where("created_by = ?", userId).
		Find(&quizzes).Error; err != nil {
		utils.ErrorResponse(c, 500, "Failed to fetch quizzes")
		return
	}

	utils.SuccessResponse(c, 200, "Quizzes fetched successfully", gin.H{"quizzes": quizzes})
}

func GetQuizByID(c *gin.Context) {
	quizID := c.Param("id")

	var quiz models.Quiz
	if err := database.DB.
		Preload("Questions.Options").
		First(&quiz, "id = ?", quizID).Error; err != nil {
		utils.ErrorResponse(c, 404, "Quiz not found")
		return
	}

	utils.SuccessResponse(c, 200, "Quiz fetched successfully", gin.H{"quiz": quiz})
}
