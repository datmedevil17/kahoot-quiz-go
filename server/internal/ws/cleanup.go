package ws

import "time"

func StartRoomTTL(hub *Hub, room *Room) {
	time.AfterFunc(30*time.Minute, func() {
		hub.RemoveRoom(room.PIN)
	})
}
