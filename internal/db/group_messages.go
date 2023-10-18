package db

import (
	"fmt"

	"github.com/dsxg666/chatroom/global/dbg"
)

type GroupMessage struct {
	Id            string
	SenderAccount string
	GroupId       string
	Message       string
	Type          string
	CreatedAt     string
}

func (p *GroupMessage) Add() {
	_, err := dbg.MySqlConn.Exec("insert group_messages set sender_account = ?, message = ?;",
		p.SenderAccount, p.Message)
	if err != nil {
		fmt.Println(err)
	}
}

func (g *GroupMessage) GetMessage() []*GroupMessage {
	var arr []*GroupMessage
	rows, err := dbg.MySqlConn.Query("select sender_account, message, created_at from group_messages where created_at >= NOW() - INTERVAL 1 DAY order by created_at asc;")
	if err != nil {
		fmt.Println("SelectAll Query err:", err)
	}
	defer rows.Close()
	for rows.Next() {
		temp := new(GroupMessage)
		err = rows.Scan(&temp.SenderAccount, &temp.Message, &temp.CreatedAt)
		if err != nil {
			fmt.Println("SelectAll Scan err:", err)
		}
		arr = append(arr, temp)
	}
	return arr
}
