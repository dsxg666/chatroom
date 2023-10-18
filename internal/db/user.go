package db

import (
	"fmt"

	"github.com/dsxg666/chatroom/global/dbg"
)

type User struct {
	Account  string
	Username string
	Password string
	Img      string
}

func (u *User) Add() {
	_, err := dbg.MySqlConn.Exec("insert user set username = ?, password = ?",
		u.Username, u.Password)
	if err != nil {
		fmt.Println(err)
	}
}

func (u *User) GetAccount() {
	stmt, err := dbg.MySqlConn.Prepare("select account from user where username = ?")
	defer stmt.Close()
	err = stmt.QueryRow(u.Username).Scan(&u.Account)
	if err != nil {
		fmt.Println(err)
	}
}

func (u *User) GetInfo() {
	stmt, err := dbg.MySqlConn.Prepare("select username, img from user where account = ?")
	defer stmt.Close()
	err = stmt.QueryRow(u.Account).Scan(&u.Username, &u.Img)
	if err != nil {
		fmt.Println(err)
	}
}

func (u *User) IsExist() bool {
	stmt, err := dbg.MySqlConn.Prepare(
		"select count(*) as count from user where username = ?")
	defer stmt.Close()
	var i int
	err = stmt.QueryRow(u.Username).Scan(&i)
	if err != nil {
		fmt.Println(err)
	}
	return i >= 1
}

func (u *User) IsExist2() bool {
	stmt, err := dbg.MySqlConn.Prepare(
		"select count(*) as count from user where account = ?")
	defer stmt.Close()
	var i int
	err = stmt.QueryRow(u.Account).Scan(&i)
	if err != nil {
		fmt.Println(err)
	}
	return i >= 1
}

func (u *User) IsCorrect() bool {
	stmt, err := dbg.MySqlConn.Prepare("select password from user where account = ?")
	defer stmt.Close()
	var passwd string
	err = stmt.QueryRow(u.Account).Scan(&passwd)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return passwd == u.Password
}

func (u *User) UpdateName() {
	_, err := dbg.MySqlConn.Exec("update user set username = ? where account = ?", u.Username, u.Account)
	if err != nil {
		fmt.Println(err)
	}
}

func (u *User) UpdateImg() {
	_, err := dbg.MySqlConn.Exec("update user set img = ? where account = ?", u.Img, u.Account)
	if err != nil {
		fmt.Println(err)
	}
}
