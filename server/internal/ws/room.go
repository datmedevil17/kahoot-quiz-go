package ws

import (
	"sync"
	"time"

	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
)

type Room struct {
	PIN     string
	HostID  string
	Clients map[string]*Client
	QuizID  string
	Hub     *Hub
	Mutex   sync.Mutex

	// Game state
	Questions []models.Question
	CurrentQ  int
	Started   bool
	Answers   map[string]int // clientID → option index
	Scores    map[string]int // clientID → total score
	StartTime time.Time
}
