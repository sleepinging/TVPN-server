package dao

import (
	"database/sql"
	// "fmt"
	"module"
)

var stmtUserAdd *sql.Stmt
var stmtUserDeleteName *sql.Stmt

func initUserDB() (err error) {
	//插入数据
	stmtUserAdd, err = db.Prepare(`INSERT INTO user("name", "password", "group") values(?,?,?)`)
	if err != nil {
		return
	}
	//删除数据
	stmtUserDeleteName, err = db.Prepare("delete from user where name=?")
	if err != nil {
		return
	}
	return
}

//AddUser 添加一个用户到数据库
func AddUser(username, password string, groupid int) (err error) {
	_, err = stmtUserAdd.Exec(username, password, groupid)
	return
}

//GetUserByName byname
func GetUserByName(username string) (user *module.User) {
	rows, err := db.Query("SELECT * FROM user where name = " + username)
	if err != nil {
		return
	}
	for rows.Next() {
		user = new(module.User)
		err = rows.Scan(&user.Id, nil, &user.Password, &user.GroupId)
		break
	}
	return
}

//DeleteUserByName delete
func DeleteUserByName(username string) (err error) {
	_, err = stmtUserDeleteName.Exec(username)
	return
}
