package ws

// Hub maintains the set of active Clients and broadcasts messages to the Clients.
type Hub struct {
	// Registered Clients.
	Clients map[*Client]bool

	ClientsMap map[string]*Client

	// Inbound messages from the Clients.
	Broadcast chan *Message

	Private chan *Message

	// Register requests from the Clients.
	Register chan *Client

	// Unregister requests from Clients.
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		ClientsMap: make(map[string]*Client),
		Broadcast:  make(chan *Message),
		Private:    make(chan *Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			h.ClientsMap[client.Host] = client
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				delete(h.ClientsMap, client.Host)
				close(client.Send)
			}
		case msgObj := <-h.Private:
			client := h.ClientsMap[msgObj.Sender]
			client.Send <- msgObj
			client2, oK := h.ClientsMap[msgObj.Receiver]
			if oK {
				client2.Send <- msgObj
			}
		case msgObj := <-h.Broadcast:
			for client := range h.Clients {
				client.Send <- msgObj
			}
		}
	}
}
