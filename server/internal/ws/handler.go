package ws

import (
	"fmt"
	"net/http"

	"github.com/datmedevil17/kahoot-quiz-go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all origins (lock later)
	},
}

func HandleWS(hub *Hub, jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := utils.ValidateToken(token, jwtSecret)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		conn, err := upgrader.Upgrade(
			c.Writer,
			c.Request,
			nil,
		)
		if err != nil {
			return
		}

		// Use claims for user info
		// Note: UserID from claims is uint, converted to string for Client
		client := &Client{
			ID:       utils.GenerateID(),
			UserID:   fmt.Sprintf("%d", claims.UserID),
			Email:    claims.Email,
			Username: claims.Email, // Default username is email, can be updated later
			Conn:     conn,
			Send:     make(chan []byte),
		}

		go writePump(client)
		readPump(client, hub)
	}
}
