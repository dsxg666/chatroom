package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dsxg666/chatroom/internal/db"
	"github.com/gorilla/websocket"
)

type Message struct {
	Message   string `json:"Message"`
	Sender    string `json:"Sender"`
	Receiver  string `json:"Receiver"`
	SenderImg string `json:"SenderImg"`
	Time      string `json:"Time"`
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	Host string
	Conn *websocket.Conn
	Send chan *Message
	Hub  *Hub
}

func (c *Client) ReadMsg() {
	defer func() {
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
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
		if msgObj.Receiver == "WorldRoom" {
			groupMessage := &db.GroupMessage{SenderAccount: msgObj.Sender, Message: msgObj.Message}
			groupMessage.Add()
			c.Hub.Broadcast <- msgObj
		} else {
			privateMsg := &db.PrivateMessage{SenderAccount: msgObj.Sender, ReceiverAccount: msgObj.Receiver, Message: msgObj.Message}
			privateMsg.Add()
			c.Hub.Private <- msgObj
		}
	}
}

func (c *Client) WriteMsg() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case msgObj, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The Hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			senderUser := &db.User{Account: msgObj.Sender}
			senderUser.GetInfo()
			msgObj.SenderImg = senderUser.Img
			msgObj.Time = time.Now().Format("2006-01-02 15:04")
			jsonMsg, err := json.Marshal(msgObj)
			if err != nil {
				fmt.Println("JSON serialization error:", err)
				return
			}
			w.Write(jsonMsg)

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
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
	// Allow collection of memory referenced by the caller by doing all work in new goroutines.
	go client.WriteMsg()
	go client.ReadMsg()
}
