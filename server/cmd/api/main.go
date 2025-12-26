package main

import (
	"log"

	"github.com/datmedevil17/kahoot-quiz-go/internal/api"
	"github.com/datmedevil17/kahoot-quiz-go/internal/config"
	"github.com/datmedevil17/kahoot-quiz-go/internal/database"
	"github.com/datmedevil17/kahoot-quiz-go/internal/services/game"
	"github.com/datmedevil17/kahoot-quiz-go/internal/ws"
)

func main() {
	// 1. Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	if err := cfg.Validate(); err != nil {
		log.Printf("Config warning: %v", err)
	}

	// 2. Connect to Database
	if err := database.Connect(cfg.DatabaseURL); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 3. Run Migrations
	if err := database.Migrate(); err != nil {
		log.Printf("Migration warning: %v", err)
	}

	// 4. Initialize Services
	db := database.GetDB()
	gameService := game.NewService(db)

	// 5. Initialize WebSocket Hub
	hub := ws.NewHub(gameService)

	// 6. Setup Router
	r := api.SetupRouter(cfg, hub)

	// 7. Start Server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
