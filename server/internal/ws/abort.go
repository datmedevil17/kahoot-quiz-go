package ws

func abortGame(room *Room, hub *Hub) {
	// Note: Caller usually holds room lock if checking condition on room, but verify broadcast.
	// If broadcast iterates clients, we need room lock.
	// If we are called from handleDisconnect, we have the lock.

	broadcast(room, "game_aborted", nil)

	// Cleanup room
	hub.RemoveRoom(room.PIN)
}
