package service

import (
	"controller"
	"fmt"
	"github.com/bitly/go-simplejson"
	"net/http"
	"strings"
)

//AddUser add user
func AddUser(w http.ResponseWriter, r *http.Request) {
	js, err := Parse2Json(r)
	if err != nil {
		fmt.Fprintln(w, GenRes(1, err.Error()))
	}
	name := js.Get("username").MustString("")
	pwd := js.Get("pwd").MustString("")
	gid := js.Get("group").MustInt(-1)
	err = controller.AddUser(name, pwd, gid)
	if err != nil {
		fmt.Fprintln(w, GenRes(1, err.Error()))
		return
	}
	fmt.Fprintln(w, GenRes(0, "添加成功"))
}

//GetUsers get user
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := controller.GetUsers()
	var jss []string
	for _, user := range users {
		ju := simplejson.New()
		ju.Set("Name", user.Name)
		ju.Set("Group", GetGroupNameByID(user.GroupID))
		b, _ := ju.MarshalJSON()
		jss = append(jss, string(b))
	}
	jstr := strings.Join(jss, ",")
	fmt.Fprintln(w, "["+jstr+"]")
}
