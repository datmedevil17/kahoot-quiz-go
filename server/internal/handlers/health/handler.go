package health

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for health check
	},
}

// HealthCheck handles HTTP GET /health
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"timestamp": time.Now(),
		"service":   "kahoot-quiz-server",
	})
}

// WSHealthCheck handles WebSocket GET /health/ws
func WSHealthCheck(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		// Upgrade handles the error response
		return
	}
	defer conn.Close()

	// Send a ping/hello message
	if err := conn.WriteJSON(gin.H{
		"status":  "ok",
		"message": "websocket health check successful",
	}); err != nil {
		return
	}

	// Wait for a moment or simple echo loop if needed?
	// For a simple health check, just connecting and reading one message or writing one message is often enough.
	// We'll read one message to acknowledge connection if client sends anything, then close.
	// Or just close after sending the hello.

	// Let's keep it open for a second then close, or read once.
	// Simple behavior: Write OK, then close.
	conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Health check passed"))
}
