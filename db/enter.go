package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	db "github.com/moonman/mbank/db/sqlc"
	"github.com/moonman/mbank/global"
)

func Init() {
	conn, err := sql.Open(global.Config.DbName, global.Config.DbSource)
	if err != nil {
		panic(err)
	}
	db.Dao = db.NewStore(conn)
}
