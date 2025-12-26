package ws

import (
	"sync"

	"github.com/datmedevil17/kahoot-quiz-go/internal/services/game"
)

type Hub struct {
	Rooms       map[string]*Room
	GameService *game.Service
	Mutex       sync.Mutex
}

func NewHub(gameService *game.Service) *Hub {
	return &Hub{
		Rooms:       make(map[string]*Room),
		GameService: gameService,
	}
}

func (h *Hub) RemoveRoom(pin string) {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()
	delete(h.Rooms, pin)
}

func (h *Hub) AddRoom(room *Room) {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()
	h.Rooms[room.PIN] = room
}
