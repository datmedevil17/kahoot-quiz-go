package database

import (
	"log"

	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
)

func Migrate() error {
	err := DB.AutoMigrate(&models.User{}, &models.BaseModel{}, &models.Quiz{}, &models.Question{}, &models.Option{}, &models.GameSession{}, &models.PlayerScore{})
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}
	log.Println("Migrations completed successfully")
	return nil
}
