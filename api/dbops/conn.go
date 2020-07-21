package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dirverName := "mysql"
	dataSourceName := "root:mim616293@tcp(localhost:3306)/video_server?charset=utf8"
	dbConn, err = sql.Open(dirverName, dataSourceName)
	if err != nil {
		panic("link mysql err: " + err.Error())
		return
	}
	err = dbConn.Ping()
	if err != nil {
		panic("ping mysql err: " + err.Error())
		return
	}
	log.Printf("link mysql successfully !")
}
