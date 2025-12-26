package models

// WSMessage is the standard envelope for all WebSocket communication
// Every message sent or received MUST follow this format
type WSMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

type JoinGamePayload struct {
	GamePIN  string `json:"game_pin"`
	Username string `json:"username"`
}