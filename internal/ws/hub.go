package ws

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

type Hub struct {
	Clients       map[*Client]bool
	ClientsMap    map[string]*Client
	Broadcast     chan *Message
	Private       chan *Message
	FriendRequest chan *Message
	Register      chan *Client
	Unregister    chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Clients:       make(map[*Client]bool),
		ClientsMap:    make(map[string]*Client),
		Broadcast:     make(chan *Message),
		Private:       make(chan *Message),
		FriendRequest: make(chan *Message),
		Register:      make(chan *Client),
		Unregister:    make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			h.ClientsMap[client.Host] = client
		case client := <-h.Unregister:
			delete(h.Clients, client)
			delete(h.ClientsMap, client.Host)
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

func (h *Hub) HeartbeatCheck() {
	for {
		for client := range h.Clients {
			fmt.Println("发送心跳给账号", client.Host)
			err := client.Conn.WriteMessage(websocket.TextMessage, []byte("heartbeat"))
			if err != nil {
				fmt.Println("发送失败，账号页面已关闭：", err)
				h.Unregister <- client
			}
		}
		time.Sleep(20 * time.Second)
	}
}
