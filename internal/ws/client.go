package ws

import (
	"encoding/json"
	"fmt"
	"github.com/dsxg666/chatroom/internal/db"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Message          string `json:"Message"`
	Sender           string `json:"Sender"`
	SenderImg        string `json:"SenderImg"`
	SenderUsername   string `json:"SenderUsername"`
	Receiver         string `json:"Receiver"`
	ReceiverImg      string `json:"ReceiverImg"`
	ReceiverUsername string `json:"ReceiverUsername"`
	Time             string `json:"Time"`
	Type             string `json:"Type"`
	InfoId           string `json:"InfoId"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	Host string
	Conn *websocket.Conn
	Send chan *Message
	Hub  *Hub
}

func (c *Client) ReadMsg() {
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		msgObj := new(Message)
		err = json.Unmarshal(message, msgObj)
		if err != nil {
			fmt.Println("反序列化失败:", err)
			return
		}
		if msgObj.Type == "Message" {
			if msgObj.Receiver == "WorldRoom" {
				groupMessage := &db.GroupMessage{SenderAccount: msgObj.Sender, Message: msgObj.Message}
				groupMessage.Add()
				c.Hub.Broadcast <- msgObj
			} else {
				privateMsg := &db.PrivateMessage{SenderAccount: msgObj.Sender, ReceiverAccount: msgObj.Receiver, Message: msgObj.Message}
				privateMsg.Add()
				c.Hub.Private <- msgObj
			}
		} else if msgObj.Type == "FriendRequest" {
			c.Hub.FriendRequest <- msgObj
		} else if msgObj.Type == "AddFriend" {
			c.Hub.Private <- msgObj
		}
	}
}

func (c *Client) WriteMsg() {
	for {
		select {
		case msgObj, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			if msgObj.Type == "Message" {
				senderUser := &db.User{Account: msgObj.Sender}
				senderUser.GetInfo()
				msgObj.SenderImg = senderUser.Img
				msgObj.SenderUsername = senderUser.Username
				msgObj.Time = time.Now().Format("2006-01-02 15:04")
			} else if msgObj.Type == "FriendRequest" {

			} else if msgObj.Type == "AddFriend" {
				senderUser := &db.User{Account: msgObj.Sender}
				senderUser.GetInfo()
				msgObj.SenderImg = senderUser.Img
				msgObj.SenderUsername = senderUser.Username
				receiverUser := &db.User{Account: msgObj.Receiver}
				receiverUser.GetInfo()
				msgObj.ReceiverImg = receiverUser.Img
				msgObj.ReceiverUsername = receiverUser.Username
			}

			jsonMsg, err := json.Marshal(msgObj)
			if err != nil {
				fmt.Println("JSON serialization error:", err)
				return
			}
			w.Write(jsonMsg)

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

// ServeWs handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{Hub: hub, Conn: conn, Send: make(chan *Message), Host: r.URL.Query().Get("account")}
	client.Hub.Register <- client
	go client.WriteMsg()
	go client.ReadMsg()
}
