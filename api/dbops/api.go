package dbops

import (
	"log"
)

func AddUserCredential(loginName, passWord string) error {
	stmt, err := dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES (?,?)")
	if err != nil {
		log.Printf("dbConn.Prepare err:%v\n", err)
		return err
	}

	_, err = stmt.Exec(loginName, passWord)
	if err != nil {
		log.Printf("stmt.Exec	 err:%v\n", err)
		return err
	}
	//必须调用
	defer stmt.Close()
	return nil
}

