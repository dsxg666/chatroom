package routers

import (
	"github.com/dsxg666/chatroom/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.Engine) {
	router := r.Group("/")
	{
		// GET
		router.GET("/", controllers.Controller{}.Index)
		router.GET("/ws", controllers.Controller{}.Ws)
		router.GET("/login", controllers.Controller{}.Login)
		router.GET("/register", controllers.Controller{}.Register)

		// POST
		router.POST("/login", controllers.Controller{}.LoginP)
		router.POST("/register", controllers.Controller{}.RegisterP)
		router.POST("/isExist", controllers.Controller{}.IsExist)
		router.POST("/updateName", controllers.Controller{}.UpdateName)
		router.POST("/updateImg", controllers.Controller{}.UpdateImg)
		router.POST("/addFriend", controllers.Controller{}.AddFriend)
		router.POST("/rejectReq", controllers.Controller{}.RejectReq)
		router.POST("/agreeReq", controllers.Controller{}.AgreeReq)
		router.POST("/getChatMsg", controllers.Controller{}.GetChatMsg)
		router.POST("/getWorldRoomMsg", controllers.Controller{}.GetWorldRoomMsg)
	}
}
