package question

import (
	"strconv"

	"github.com/datmedevil17/kahoot-quiz-go/internal/database"
	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
	"github.com/datmedevil17/kahoot-quiz-go/internal/utils"
	"github.com/gin-gonic/gin"
)

func AddQuestion(c *gin.Context) {
	quizID := c.Param("quizId")
	userID, _ := c.Get("userID")

	_, err := utils.ValidateQuizOwnership(quizID, uint(userID.(uint)))
	if err != nil {
		utils.ErrorResponse(c, 403, err.Error())
		return
	}

	var req CreateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	if req.Answer < 0 || req.Answer >= len(req.Options) {
		utils.ErrorResponse(c, 400, "Answer index out of range")
		return
	}

	question := models.Question{
		QuizID:    quizID,
		Text:      req.Text,
		Answer:    req.Answer,
		TimeLimit: req.TimeLimit,
	}

	if err := database.DB.Create(&question).Error; err != nil {
		utils.ErrorResponse(c, 500, "Failed to create question")
		return
	}

	var options []models.Option
	for _, opt := range req.Options {
		options = append(options, models.Option{
			QuestionID: strconv.FormatUint(uint64(question.ID), 10),
			Text:       opt,
		})
	}

	if err := database.DB.Create(&options).Error; err != nil {
		utils.ErrorResponse(c, 500, "Failed to create options")
		return
	}

	utils.SuccessResponse(c, 201, "Question created successfully", gin.H{
		"question": question,
		"options":  options,
	})
}

func GetQuizQuestions(c *gin.Context) {
	quizID := c.Param("quizId")

	var questions []models.Question
	if err := database.DB.
		Where("quiz_id = ?", quizID).
		Preload("Options").
		Find(&questions).Error; err != nil {
		utils.ErrorResponse(c, 500, "Failed to fetch questions")
		return
	}

	utils.SuccessResponse(c, 200, "Questions fetched successfully", gin.H{"questions": questions})
}
