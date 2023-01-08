package db

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbName = "mysql"
	dbConn = "root:root@tcp(127.0.0.1:3306)/bank?parseTime=true&loc=Local"
)

var testQueries Store

//作为一个包内test文件的开始
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbName, dbConn)
	if err != nil {
		panic(err)
	}

	q := &Queries{
		db: conn,
	}
	testQueries = &SqlStore{
		Queries: q,
		db:      conn,
	}
	os.Exit(m.Run()) //获取测试的返回值
}
