package ws

func (r *Room) addClient(c *Client) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	r.Clients[c.ID] = c
}

func (r *Room) removeClient(c *Client) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Clients, c.ID)
}
