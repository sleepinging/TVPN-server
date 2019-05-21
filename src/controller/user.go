package controller

import (
	"dao"
	"errors"
	"module"
)

//AddUser add user
func AddUser(name, pwd string, groupid int) (err error) {
	if name == "" || pwd == "" || groupid == -1 {
		return errors.New("参数错误")
	}
	err = dao.AddUser(name, pwd, groupid)
	return
}

//GetUsers get user
func GetUsers() (users []*module.User) {
	users, _ = dao.GetUsers()
	return
}

//CheckLogin check login
func CheckLogin(name, pwd string) (ok bool) {
	user, err := dao.GetUserByName(name)
	if err != nil {
		return
	}
	if user == nil {
		return
	}
	if user.Password != pwd {
		return
	}
	ok = true
	return
}
