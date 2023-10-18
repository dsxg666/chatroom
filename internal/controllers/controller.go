package controllers

import (
	"fmt"
	"net/http"
	"path"

	"github.com/dsxg666/chatroom/global/wsg"
	"github.com/dsxg666/chatroom/internal/db"
	"github.com/dsxg666/chatroom/internal/ws"
	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (t Controller) Index(c *gin.Context) {
	acc := c.Query("account")
	user := &db.User{Account: acc}
	user.GetInfo()
	info := &db.Info{}
	userRelationship := &db.UserRelationships{UserAccount: acc}
	friends := userRelationship.GetFriend()
	var friendsInfo []*db.User
	for i := 0; i < len(friends); i++ {
		userTemp := &db.User{Account: friends[i].FriendAccount}
		userTemp.GetInfo()
		friendsInfo = append(friendsInfo, userTemp)
	}
	c.HTML(http.StatusOK, "main/index.html", gin.H{
		"img":       user.Img,
		"username":  user.Username,
		"infoArr":   info.GetAll(acc),
		"friendArr": friendsInfo,
	})
}

func (t Controller) Ws(c *gin.Context) {
	ws.ServeWs(wsg.Hub, c.Writer, c.Request)
}

func (t Controller) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "main/login.html", nil)
}

func (t Controller) LoginP(c *gin.Context) {
	acc := c.PostForm("account")
	user := &db.User{Account: acc, Password: c.PostForm("password")}
	_, ok := wsg.Hub.ClientsMap[acc]
	if user.IsCorrect() {
		if ok {
			c.JSON(http.StatusOK, gin.H{
				"status": "2",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "1",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "0",
		})
	}
}

func (t Controller) Quit(c *gin.Context) {
	acc := c.PostForm("account")
	client, ok := wsg.Hub.ClientsMap[acc]
	if ok {
		client.Hub.Unregister <- client
	}
	c.JSON(http.StatusOK, nil)
}

func (t Controller) Register(c *gin.Context) {
	c.HTML(http.StatusOK, "main/register.html", nil)
}

func (t Controller) RegisterP(c *gin.Context) {
	user := &db.User{Username: c.PostForm("username"), Password: c.PostForm("password")}
	user.Add()
	user.GetAccount()
	c.HTML(http.StatusOK, "main/account.html", gin.H{
		"account": user.Account,
	})
}

func (t Controller) IsExist(c *gin.Context) {
	user := &db.User{Username: c.PostForm("username")}
	if user.IsExist() {
		c.JSON(http.StatusOK, gin.H{
			"status": "0",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "1",
		})
	}
}

func (t Controller) UpdateName(c *gin.Context) {
	user := &db.User{Account: c.PostForm("account"), Username: c.PostForm("username")}
	if user.IsExist() {
		c.JSON(http.StatusOK, gin.H{
			"status": "0",
		})
	} else {
		user.UpdateName()
		c.JSON(http.StatusOK, gin.H{
			"status": "1",
		})
	}
}

func (t Controller) UpdateImg(c *gin.Context) {
	file, _ := c.FormFile("file")
	acc := c.PostForm("account")
	filename := acc + ".jpg"
	user := &db.User{Account: acc, Img: "./static/img/avatar/" + filename}
	user.UpdateImg()
	dst := path.Join("./static/img/avatar", filename)
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, nil)
}

func (t Controller) AddFriend(c *gin.Context) {
	receiver := c.PostForm("receiver")
	sender := c.PostForm("account")
	if receiver == sender {
		c.JSON(http.StatusOK, gin.H{
			"status": "2",
		})
	} else {
		user := &db.User{Account: receiver}
		if user.IsExist2() {
			info := &db.Info{
				SenderAccount:   sender,
				ReceiverAccount: receiver,
				Type:            "0",
				Finish:          "0",
			}
			userRelationship := &db.UserRelationships{UserAccount: sender, FriendAccount: receiver}
			if userRelationship.IsExist() {
				c.JSON(http.StatusOK, gin.H{
					"status": "3",
				})
			} else {
				if info.IsFinish() {
					info.Add()
					c.JSON(http.StatusOK, gin.H{
						"status": "1",
					})
				} else {
					c.JSON(http.StatusOK, gin.H{
						"status": "-1",
					})
				}
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "0",
			})
		}
	}
}

func (t Controller) RejectReq(c *gin.Context) {
	infoId := c.PostForm("infoId")
	info := &db.Info{Id: infoId}
	info.FinishReq()
	c.JSON(http.StatusOK, nil)
}

func (t Controller) AgreeReq(c *gin.Context) {
	infoId := c.PostForm("infoId")
	sender := c.PostForm("sender")
	senderUser := &db.User{Account: sender}
	senderUser.GetInfo()
	receiver := c.PostForm("receiver")
	receiverUser := &db.User{Account: receiver}
	receiverUser.GetInfo()
	userRelationship := &db.UserRelationships{UserAccount: sender, FriendAccount: receiver}
	userRelationship.Add()
	info := &db.Info{Id: infoId}
	info.FinishReq()
	privateMsg := &db.PrivateMessage{SenderAccount: sender, ReceiverAccount: receiver, Message: "我是" + senderUser.Username + "，我们是好友了，一起来聊天吧。[初始化消息]"}
	privateMsg.Add()
	privateMsg2 := &db.PrivateMessage{SenderAccount: receiver, ReceiverAccount: sender, Message: "我是" + receiverUser.Username + "，很高兴和你成为好友，一起来聊天吧。[初始化消息]"}
	privateMsg2.Add()
	c.JSON(http.StatusOK, nil)
}

func (t Controller) GetChatMsg(c *gin.Context) {
	sender := c.PostForm("sender")
	senderUser := &db.User{Account: sender}
	senderUser.GetInfo()
	receiver := c.PostForm("receiver")
	receiverUser := &db.User{Account: receiver}
	receiverUser.GetInfo()
	privateMessage := &db.PrivateMessage{SenderAccount: sender, ReceiverAccount: receiver}
	c.JSON(http.StatusOK, gin.H{
		"msgs":             privateMessage.GetMessage(),
		"senderImg":        senderUser.Img,
		"receiverImg":      receiverUser.Img,
		"receiverUsername": receiverUser.Username,
	})
}

func (t Controller) GetWorldRoomMsg(c *gin.Context) {
	groupMessage := &db.GroupMessage{}
	msgArr := groupMessage.GetMessage()
	msgArrLen := len(msgArr)
	var imgArr []string
	for i := 0; i < msgArrLen; i++ {
		senderUser := &db.User{Account: msgArr[i].SenderAccount}
		senderUser.GetInfo()
		imgArr = append(imgArr, senderUser.Img)
	}
	c.JSON(http.StatusOK, gin.H{
		"msgWorldRoom": msgArr,
		"imgWorldRoom": imgArr,
	})
}
