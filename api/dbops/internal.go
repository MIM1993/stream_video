package dbops

import (
	"strconv"
	"log"
)
func InsertSession(userName string, ttl int64, uuid string) error {
	stmt, err := dbConn.Prepare("INSERT INTO sessions (session_id,TTL,login_name) values (?,?,?)")
	if err != nil {
		log.Printf("dbConn.Prepare err:%v\n", err)
		return err
	}
	ttlStr := strconv.FormatInt(ttl, 10)

	_,err = stmt.Exec(uuid,ttlStr,userName)
	if err != nil {
		log.Printf("stmt.Exec	 err:%v\n", err)
		return err
	}
	//必须调用
	defer stmt.Close()
	return nil
}