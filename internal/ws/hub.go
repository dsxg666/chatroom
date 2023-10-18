package ws

type Hub struct {
	Clients map[*Client]bool

	ClientsMap map[string]*Client

	Broadcast chan *Message

	Private chan *Message

	Register chan *Client

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
