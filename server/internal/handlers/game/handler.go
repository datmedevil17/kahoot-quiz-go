package game

import (
	"fmt"
	"net/http"

	"github.com/datmedevil17/kahoot-quiz-go/internal/services/game"
	"github.com/datmedevil17/kahoot-quiz-go/internal/utils"
	"github.com/datmedevil17/kahoot-quiz-go/internal/ws"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *game.Service
}

func NewHandler(service *game.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateGame(hub *ws.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")

		pin := utils.GeneratePIN()

		room := &ws.Room{
			PIN:     pin,
			HostID:  fmt.Sprintf("%v", userID),
			Clients: make(map[string]*ws.Client),
			Hub:     hub,
		}

		hub.AddRoom(room)
		ws.StartRoomTTL(hub, room)

		utils.SuccessResponse(c, http.StatusCreated, "Game created successfully", gin.H{
			"game_pin": pin,
		})
	}
}
