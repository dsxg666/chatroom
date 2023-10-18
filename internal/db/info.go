package db

import (
	"fmt"

	"github.com/dsxg666/chatroom/global/dbg"
)

type Info struct {
	Id              string
	SenderAccount   string
	ReceiverAccount string
	Type            string
	Finish          string
}

func (i *Info) Add() {
	_, err := dbg.MySqlConn.Exec("insert info set sender_account = ?, receiver_account = ?, `type` = ?, `finish` = ?",
		i.SenderAccount, i.ReceiverAccount, i.Type, i.Finish)
	if err != nil {
		fmt.Println(err)
	}
}

func (i *Info) GetAll(receiver string) []*Info {
	var arr []*Info
	rows, err := dbg.MySqlConn.Query("select * from info where receiver_account = ? and `finish` = 0 order by id desc", receiver)
	if err != nil {
		fmt.Println("SelectAll Query err:", err)
	}
	defer rows.Close()
	for rows.Next() {
		temp := new(Info)
		err = rows.Scan(&temp.Id, &temp.SenderAccount, &temp.ReceiverAccount, &temp.Type, &temp.Finish)
		if err != nil {
			fmt.Println("SelectAll Scan err:", err)
		}
		arr = append(arr, temp)
	}
	return arr
}

func (i *Info) FinishReq() {
	_, err := dbg.MySqlConn.Exec("update info set `finish` = 1 where id = ?", i.Id)
	if err != nil {
		fmt.Println(err)
	}
}

func (i *Info) IsFinish() bool {
	stmt, err := dbg.MySqlConn.Prepare(
		"select count(*) as count from info where sender_account = ? and receiver_account = ? and finish = '0'")
	defer stmt.Close()
	var t int
	err = stmt.QueryRow(i.SenderAccount, i.ReceiverAccount).Scan(&t)
	if err != nil {
		fmt.Println(err)
	}
	return t < 1
}
