package dao

import (
	"config"
	"database/sql"
	"path"
)

var db *sql.DB

//InitDB init db
func InitDB() (err error) {
	db, err = sql.Open("sqlite3", path.Join(config.WorkPath, "data.db"))
	if err != nil {
		return
	}
	err = initUserDB()
	if err != nil {
		return
	}
	err = initGroupDB()
	if err != nil {
		return
	}
	err = initRecordDB()
	if err != nil {
		return
	}
	err = initAccessDB()
	if err != nil {
		return
	}
	return
}
