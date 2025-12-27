package question

import (
	"github.com/datmedevil17/kahoot-quiz-go/internal/services/question"
	"github.com/datmedevil17/kahoot-quiz-go/internal/services/quiz"
	"github.com/datmedevil17/kahoot-quiz-go/internal/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	questionService *question.Service
	quizService     *quiz.Service
}

func NewHandler(questionService *question.Service, quizService *quiz.Service) *Handler {
	return &Handler{
		questionService: questionService,
		quizService:     quizService,
	}
}

func (h *Handler) AddQuestion(c *gin.Context) {
	quizID := c.Param("id")
	userID, _ := c.Get("userID")

	if err := h.quizService.ValidateOwnership(quizID, uint(userID.(uint))); err != nil {
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

	createdQuestion, options, err := h.questionService.CreateQuestion(quizID, req.Text, req.Answer, req.TimeLimit, req.Options)
	if err != nil {
		utils.ErrorResponse(c, 500, "Failed to create question")
		return
	}

	utils.SuccessResponse(c, 201, "Question created successfully", gin.H{
		"question": createdQuestion,
		"options":  options,
	})
}

func (h *Handler) GetQuizQuestions(c *gin.Context) {
	quizID := c.Param("id")

	questions, err := h.questionService.GetQuestionsByQuizID(quizID)
	if err != nil {
		utils.ErrorResponse(c, 500, "Failed to fetch questions")
		return
	}

	utils.SuccessResponse(c, 200, "Questions fetched successfully", gin.H{"questions": questions})
}
