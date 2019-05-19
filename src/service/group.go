package service

import (
	"dao"
)

//GetGroupNameById get name
func GetGroupNameByID(id int) (name string) {
	group, _ := dao.GetGroupByID(id)
	if group != nil {
		name = group.Name
	}
	return
}
