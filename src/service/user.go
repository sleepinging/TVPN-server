package service

import (
	"dao"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bitly/go-simplejson"
)

//AddUser add user
func AddUser(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "get data failed")
		return
	}
	js, err := simplejson.NewJson(data)
	if err != nil {
		fmt.Fprintln(w, "get json data failed")
		return
	}
	name := js.Get("username").MustString()
	pwd := js.Get("pwd").MustString()
	gid := js.Get("group").MustInt()
	err = dao.AddUser(name, pwd, gid)
	if err != nil {
		fmt.Fprintln(w, GenRes(1, err.Error()))
		return
	}
	fmt.Fprintln(w, GenRes(0, "添加成功"))
}
