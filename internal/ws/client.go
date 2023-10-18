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
			if msgObj.Sender == c.Host {
				c.Hub.ClientTime[c] = time.Now()
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

func (c *Client) CheckClientActivity() {
	for {
		// 检查客户端的最后活动时间
		lastActiveTime, ok := c.Hub.ClientTime[c]
		if !ok {
			return
		}
		// 如果客户端在一段时间内没有发出任何消息，将其断开连接
		if time.Since(lastActiveTime) > 300*time.Second {
			fmt.Printf("用户长时间未活动，断开连接: %s\n", c.Conn.RemoteAddr())
			c.Hub.Unregister <- c
			return
		}

		time.Sleep(360 * time.Second) // 等待一段时间再次检查
	}
}

func (c *Client) CheckHeartbeat() {
	for {
		// 定时发送心跳消息给所有连接的客户端
		for client := range c.Hub.Clients {
			err := client.Conn.WriteMessage(websocket.TextMessage, []byte("heartbeat"))
			if err != nil {
				log.Println(err)
			}
		}
		time.Sleep(10 * time.Second)
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
	go client.CheckClientActivity()
	go client.CheckHeartbeat()
}
