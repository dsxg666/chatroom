package db

import (
	"fmt"
	"github.com/dsxg666/chatroom/global/dbg"
)

type UserRelationships struct {
	Id            string
	UserAccount   string
	FriendAccount string
	Status        string
	CreateAt      string
}

func (u *UserRelationships) Add() {
	_, err := dbg.MySqlConn.Exec("insert user_relationships set user_account = ?, friend_account = ?, status = 1",
		u.UserAccount, u.FriendAccount)
	if err != nil {
		fmt.Println(err)
	}
	_, err = dbg.MySqlConn.Exec("insert user_relationships set user_account = ?, friend_account = ?, status = 1",
		u.FriendAccount, u.UserAccount)
	if err != nil {
		fmt.Println(err)
	}
}
func (u *UserRelationships) IsExist() bool {
	stmt, err := dbg.MySqlConn.Prepare(
		"select count(*) as count from user_relationships where user_account = ? and friend_account = ? and status = 1 or user_account = ? and friend_account = ? and status = 1")
	defer stmt.Close()
	var i int
	err = stmt.QueryRow(u.UserAccount, u.FriendAccount, u.FriendAccount, u.UserAccount).Scan(&i)
	if err != nil {
		fmt.Println(err)
	}
	return i >= 1
}

func (u *UserRelationships) GetFriend() []*UserRelationships {
	var arr []*UserRelationships
	rows, err := dbg.MySqlConn.Query("select friend_account from user_relationships where user_account = ? and status = 1 order by friend_account desc;", u.UserAccount)
	if err != nil {
		fmt.Println("SelectAll Query err:", err)
	}
	defer rows.Close()
	for rows.Next() {
		temp := new(UserRelationships)
		err = rows.Scan(&temp.FriendAccount)
		if err != nil {
			fmt.Println("SelectAll Scan err:", err)
		}
		arr = append(arr, temp)
	}
	return arr
}
