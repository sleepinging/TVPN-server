package view

import (
	"client"
	"fmt"
	"net/http"
	"path"
	"service"
	"strings"

	"config"
)

//获取在线用户
func GetOnlineUserHandler(w http.ResponseWriter, r *http.Request) {
	// uid := r.PostFormValue("uid")
	cs := client.GetAllOnlineClient(10)
	// fmt.Println("get")
	// client.PrintOnlineClientMap(-1)
	strs := make([]string, len(cs))
	for i, v := range cs {
		strs[i] = v.ToJson()
	}
	str := strings.Join(strs, ",")
	str = "[" + str + "]"
	// data, _ := json.Marshal(cs)
	// fmt.Fprintln(w, string(data))
	// str := `[
	// 	{"Name":"twt","IP":"192.168.10.100","MAC":"00:02:03:04:05:07","OnlineTime":"2019-05-01 12:00:00"},
	// 	{"Name":"twt-centos","IP":"192.168.10.98","MAC":"00:02:03:04:09:08","OnlineTime":"2019-05-01 12:00:10"}
	// ]`
	fmt.Fprintln(w, str)
}

//获取所有用户
func GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	str := `[
		{"Name":"twt","Group":"admin"},
		{"Name":"twt-centos","Group":"admin"},
		{"Name":"ldl","Group":"Boss"},
		{"Name":"马化腾","Group":"guest"},
		{"Name":"雷军","Group":"guest"}
	]`
	fmt.Fprintln(w, str)
}

//获取用户组
func GetGroupHandler(w http.ResponseWriter, r *http.Request) {
	// str := `[
	// 	{"id":"1","Name":"admin","Network":"192.168.10.1/24","MAC":"*","up_speed_limit":123,"down_speed_limt":1234,"def_visit":1,"def_visited":0},
	// 	{"id":"2","Name":"Boss","Network":"192.168.10.2/24","MAC":"*","up_speed_limit":123,"down_speed_limt":1234,"def_visit":1,"def_visited":0},
	// 	{"id":"3","Name":"guest","Network":"192.168.10.3/24","MAC":"*","up_speed_limit":123,"down_speed_limt":1234,"def_visit":0,"def_visited":1}
	// ]`
	str := `[
		{"id":"1","Name":"user","Network":"192.168.10.1/24","MAC":"*","up_speed_limit":123,"down_speed_limt":1234,"def_visit":1,"def_visited":0}
	]`
	fmt.Fprintln(w, str)
}

//获取权限数据
func GetAccessHandler(w http.ResponseWriter, r *http.Request) {
	str := `[
		{"Name":"admin","Group1":"admin","Group2":"Boss","Rule":"1"},
		{"Name":"test","Group1":"guest","Group2":"Boss","Rule":"0"}
	]`
	fmt.Fprintln(w, str)
}

//获取记录数据
func GetRecordHandler(w http.ResponseWriter, r *http.Request) {
	str := `[
		{"Name":"twt","Online_Time":"2019-04-25 12:00:00","Off_Time":"2019-04-25 13:00:00","IP":"192.168.10.100","MAC":"02:03:04:05:06:07","UP_Flow":1234,"DOWN_Flow":123456},
		{"Name":"ldl","Online_Time":"2019-04-26 12:00:00","Off_Time":"2019-04-26 13:00:00","IP":"192.168.10.98","MAC":"02:03:04:05:06:06","UP_Flow":1234,"DOWN_Flow":123456},
		{"Name":"twt-centos","Online_Time":"2019-04-27 12:00:00","Off_Time":"2019-04-27 13:00:00","IP":"192.168.10.97","MAC":"02:03:04:05:06:05","UP_Flow":1234,"DOWN_Flow":123456}
	]`
	fmt.Fprintln(w, str)
}

func StartHTTPServer(port int) {

	http.HandleFunc("/getonlineuser", GetOnlineUserHandler)
	http.HandleFunc("/getalluser", GetAllUserHandler)
	http.HandleFunc("/getgroup", GetGroupHandler)
	http.HandleFunc("/getaccess", GetAccessHandler)
	http.HandleFunc("/getrecord", GetRecordHandler)

	//查看在线用户
	http.HandleFunc("/onlineuser/get", GetRecordHandler)
	//下线在线用户
	http.HandleFunc("/onlineuser/offline", GetRecordHandler)

	//增加用户
	http.HandleFunc("/user/add", service.AddUser)
	//删除用户
	http.HandleFunc("/user/delete", GetRecordHandler)
	//查询用户
	http.HandleFunc("/user/get", service.GetUsers)
	//修改用户
	http.HandleFunc("/user/edit", GetRecordHandler)

	//增加用户组
	http.HandleFunc("/group/add", GetRecordHandler)
	//删除用户组
	http.HandleFunc("/group/delete", GetRecordHandler)
	//修改用户组
	http.HandleFunc("/group/edit", GetRecordHandler)
	//查询用户组
	http.HandleFunc("/group/get", GetRecordHandler)

	//增加用户组权限
	http.HandleFunc("/access/add", GetRecordHandler)
	//删除用户组权限
	http.HandleFunc("/access/delete", GetRecordHandler)
	//修改用户组权限
	http.HandleFunc("/access/edit", GetRecordHandler)
	//查询用户组权限
	http.HandleFunc("/access/get", GetRecordHandler)

	//查询记录
	http.HandleFunc("/record/get", GetRecordHandler)

	http.Handle("/", http.FileServer(http.Dir(path.Join(config.WorkPath, "html"))))

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}
