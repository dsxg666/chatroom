package dbg

import (
	"database/sql"

	"github.com/dsxg666/chatroom/pkg/setting"
)

var (
	DatabaseSetting *setting.DatabaseSettingS
	MySqlConn       *sql.DB
)
