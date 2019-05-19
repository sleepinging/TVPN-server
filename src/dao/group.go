package dao

import (
	"database/sql"
	"module"
)

var stmtGroupAdd *sql.Stmt
var stmtGroupDeleteName *sql.Stmt

func initGroupDB() (err error) {
	//插入数据
	stmtGroupAdd, err = db.Prepare(`INSERT INTO "group" ("name", "network", "mac", "up_speed_limit", "down_speed_limt", "def_visit") VALUES (?,?,?,?,?,?)`)
	if err != nil {
		return
	}
	//删除数据
	stmtGroupDeleteName, err = db.Prepare("delete from 'group' where name=?")
	if err != nil {
		return
	}
	return
}

//AddGroup addgroup
func AddGroup(name, betwork, mac string, upSpeedLimit, downSpeedLimt, defVisit int) (err error) {

	return
}

//GetGroups get all groups
func GetGroups() (groups []*module.Group) {
	return
}

//GetGroupByID get group by id
func GetGroupByID(id int) (group *module.Group, err error) {
	rows, err := db.Query("SELECT * FROM 'group'")
	if err != nil {
		return
	}
	for rows.Next() {
		group = new(module.Group)
		err = rows.Scan(&group.ID, &group.Name, &group.NetWork, &group.MAC, &group.UpSpeedLimit, &group.DownSpeedLimit, &group.DefVisit)
		break
	}
	return
}
