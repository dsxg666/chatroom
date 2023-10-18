package db

import (
	"database/sql"
	"fmt"

	"github.com/dsxg666/chatroom/global/dbg"
	"github.com/dsxg666/chatroom/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
)

func initMysql(dbSetting *setting.DatabaseSettingS) *sql.DB {
	dbConn, err := sql.Open(dbSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t",
		dbSetting.UserName,
		dbSetting.Password,
		dbSetting.Host,
		dbSetting.DBName,
		dbSetting.Charset,
		dbSetting.ParseTime,
	))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// 最大连接数
	dbConn.SetMaxOpenConns(dbg.DatabaseSetting.MaxOpenConns)
	// 闲置连接数
	dbConn.SetMaxIdleConns(dbg.DatabaseSetting.MaxIdleConns)
	return dbConn
}

func InitDB(dbSetting *setting.DatabaseSettingS) *sql.DB {
	return initMysql(dbSetting)
}
