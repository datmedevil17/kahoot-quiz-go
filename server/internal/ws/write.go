package ws

func writePump(client *Client) {
	for msg := range client.Send {
		client.Conn.WriteMessage(1, msg)
	}
}
