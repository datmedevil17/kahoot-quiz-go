package quiz

import (
	"net/http"

	"github.com/datmedevil17/kahoot-quiz-go/internal/services/quiz"
	"github.com/datmedevil17/kahoot-quiz-go/internal/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *quiz.Service
}

func NewHandler(service *quiz.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateQuiz(c *gin.Context) {
	var req CreateQuizRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}
	userId, _ := c.Get("userID")

	createdQuiz, err := h.service.CreateQuiz(req.Title, req.Description, uint(userId.(uint)))
	if err != nil {
		utils.ErrorResponse(c, 500, "Failed to create quiz")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Quiz created successfully", gin.H{"quiz": createdQuiz})
}

func (h *Handler) GetMyQuizzes(c *gin.Context) {
	userId, _ := c.Get("userID")

	quizzes, err := h.service.GetQuizzesByUserID(uint(userId.(uint)))
	if err != nil {
		utils.ErrorResponse(c, 500, "Failed to fetch quizzes")
		return
	}

	utils.SuccessResponse(c, 200, "Quizzes fetched successfully", gin.H{"quizzes": quizzes})
}

func (h *Handler) GetQuizByID(c *gin.Context) {
	quizID := c.Param("id")

	quiz, err := h.service.GetQuizByID(quizID)
	if err != nil {
		utils.ErrorResponse(c, 404, "Quiz not found")
		return
	}

	utils.SuccessResponse(c, 200, "Quiz fetched successfully", gin.H{"quiz": quiz})
}
