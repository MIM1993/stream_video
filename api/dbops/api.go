package dbops

import (
	"database/sql"
	"log"
)

func AddUserCredential(loginName, passWord string) error {
	stmt, err := dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES (?,?)")
	defer stmt.Close()
	if err != nil {
		log.Printf("dbConn.Prepare err:%v\n", err)
		return err
	}

	_, err = stmt.Exec(loginName, passWord)
	if err != nil {
		log.Printf("stmt.Exec	 err:%v\n", err)
		return err
	}
	return nil
}

func GetUserPwd(loginName string) (string, error) {
	stmt, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	defer stmt.Close()
	if err != nil {
		log.Printf("dbConn.Prepare err:%v\n", err)
		return "", err
	}

	var pwd string
	err = stmt.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	return pwd, err
}

func DeleteUser(loginName string) error {

	return nil
}

func GetUser(loginName string) {

}
