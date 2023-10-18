package ws

type Hub struct {
	Clients map[*Client]bool

	ClientsMap map[string]*Client

	Broadcast chan *Message

	Private chan *Message

	FriendRequest chan *Message

	Register chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Clients:       make(map[*Client]bool),
		ClientsMap:    make(map[string]*Client),
		Broadcast:     make(chan *Message),
		Private:       make(chan *Message),
		FriendRequest: make(chan *Message),
		Register:      make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			h.ClientsMap[client.Host] = client
		case msgObj := <-h.Private:
			client, ok := h.ClientsMap[msgObj.Sender]
			if ok {
				client.Send <- msgObj
			}
			client2, oK2 := h.ClientsMap[msgObj.Receiver]
			if oK2 {
				client2.Send <- msgObj
			}
		case msgObj := <-h.Broadcast:
			for client := range h.Clients {
				client.Send <- msgObj
			}
		case msgObj := <-h.FriendRequest:
			client, ok := h.ClientsMap[msgObj.Receiver]
			if ok {
				client.Send <- msgObj
			}
		}
	}
}
