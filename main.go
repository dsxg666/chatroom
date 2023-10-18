package main

import (
	"fmt"

	"github.com/dsxg666/chatroom/global/dbg"
	"github.com/dsxg666/chatroom/global/wsg"
	"github.com/dsxg666/chatroom/internal/db"
	"github.com/dsxg666/chatroom/internal/routers"
	"github.com/dsxg666/chatroom/internal/ws"
	"github.com/dsxg666/chatroom/pkg/setting"
	"github.com/gin-gonic/gin"
)

func init() {
	wsg.Hub = ws.NewHub()
	go wsg.Hub.Run()

	err := setupSetting()
	if err != nil {
		fmt.Println(err)
	}
	dbg.MySqlConn = db.InitDB(dbg.DatabaseSetting)
}

func main() {
	// gin.SetMode("release")
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/*")
	routers.RouterInit(r)
	r.Run()
}

func setupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &dbg.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
