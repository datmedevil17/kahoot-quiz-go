package api

import (
	"github.com/datmedevil17/kahoot-quiz-go/internal/config"
	"github.com/datmedevil17/kahoot-quiz-go/internal/database"
	"github.com/datmedevil17/kahoot-quiz-go/internal/handlers/game"
	"github.com/datmedevil17/kahoot-quiz-go/internal/handlers/health"
	"github.com/datmedevil17/kahoot-quiz-go/internal/handlers/question"
	"github.com/datmedevil17/kahoot-quiz-go/internal/handlers/quiz"
	"github.com/datmedevil17/kahoot-quiz-go/internal/handlers/user"
	"github.com/datmedevil17/kahoot-quiz-go/internal/middleware"
	questionService "github.com/datmedevil17/kahoot-quiz-go/internal/services/question"
	quizService "github.com/datmedevil17/kahoot-quiz-go/internal/services/quiz"
	userService "github.com/datmedevil17/kahoot-quiz-go/internal/services/user"
	"github.com/datmedevil17/kahoot-quiz-go/internal/ws"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config, hub *ws.Hub) *gin.Engine {
	r := gin.Default()

	// Initialize Services
	db := database.GetDB()
	userSvc := userService.NewService(db)
	quizSvc := quizService.NewService(db)
	questionSvc := questionService.NewService(db)

	// Initialize Handlers
	userHandler := user.NewHandler(userSvc, cfg.JWTSecret)
	gameHandler := game.NewHandler()
	quizHandler := quiz.NewHandler(quizSvc)
	questionHandler := question.NewHandler(questionSvc, quizSvc)

	// Auth Routes
	auth := r.Group("/auth")
	{
		auth.POST("/signup", userHandler.SignUp)
		auth.POST("/login", userHandler.SignIn)
	}

	// User Routes
	users := r.Group("/users")
	{
		users.GET("/:id", userHandler.GetUserById)
	}

	// Protected Routes
	protected := r.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		// User
		protected.GET("/users/me", userHandler.GetCurrentUserId)

		// Quiz
		protected.POST("/quizzes", quizHandler.CreateQuiz)
		protected.GET("/quizzes", quizHandler.GetMyQuizzes)
		protected.GET("/quizzes/:id", quizHandler.GetQuizByID)

		// Questions
		protected.POST("/quizzes/:id/questions", questionHandler.AddQuestion)
		protected.GET("/quizzes/:id/questions", questionHandler.GetQuizQuestions)

		// Game
		protected.POST("/games", gameHandler.CreateGame(hub))

		// WS
		protected.GET("/ws", ws.HandleWS(hub, cfg.JWTSecret))
	}

	// Health Routes
	r.GET("/health", health.HealthCheck)
	r.GET("/health/ws", health.WSHealthCheck)

	return r
}
