package db

import (
	"fmt"

	"github.com/dsxg666/chatroom/global/dbg"
)

type PrivateMessage struct {
	Id              string `json:"Id"`
	SenderAccount   string `json:"SenderAccount"`
	ReceiverAccount string `json:"ReceiverAccount"`
	Message         string `json:"Message"`
	CreatedAt       string `json:"CreatedAt"`
}

func (p *PrivateMessage) Add() {
	_, err := dbg.MySqlConn.Exec("insert private_messages set sender_account = ?, receiver_account = ?, message = ?",
		p.SenderAccount, p.ReceiverAccount, p.Message)
	if err != nil {
		fmt.Println(err)
	}
}

func (p *PrivateMessage) GetMessage() []*PrivateMessage {
	var arr []*PrivateMessage
	rows, err := dbg.MySqlConn.Query("select * from private_messages where (sender_account = ? and receiver_account = ? or sender_account = ? and receiver_account = ?) and created_at >= NOW() - INTERVAL 1 DAY order by created_at asc;",
		p.SenderAccount, p.ReceiverAccount, p.ReceiverAccount, p.SenderAccount)
	if err != nil {
		fmt.Println("SelectAll Query err:", err)
	}
	defer rows.Close()
	for rows.Next() {
		temp := new(PrivateMessage)
		err = rows.Scan(&temp.Id, &temp.SenderAccount, &temp.ReceiverAccount, &temp.Message, &temp.CreatedAt)
		if err != nil {
			fmt.Println("SelectAll Scan err:", err)
		}
		arr = append(arr, temp)
	}
	return arr
}
