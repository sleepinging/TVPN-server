package view

import (
	"fmt"
	"net/http"
)

//获取在线用户
func GetOnlineUserHandler(w http.ResponseWriter, r *http.Request) {
	// cs := client.GetAllClient(10)
	// data, _ := json.Marshal(cs)
	// fmt.Fprintln(w, string(data))
	str := `[
		{"Name":"twt","IP":"192.168.10.100","MAC":"00:02:03:04:05:07","OnlineTime":"2019-05-01 12:00:00"},
		{"Name":"twt-centos","IP":"192.168.10.98","MAC":"00:02:03:04:09:08","OnlineTime":"2019-05-01 12:00:10"}
	]`
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
	str := `[
		{"Name":"admin","Network":"192.168.10.1/24","MAC":"*","up_speed_limit":123,"down_speed_limt":1234,"def_visit":1,"def_visited":0},
		{"Name":"Boss","Network":"192.168.10.2/24","MAC":"*","up_speed_limit":123,"down_speed_limt":1234,"def_visit":1,"def_visited":0},
		{"Name":"guest","Network":"192.168.10.3/24","MAC":"*","up_speed_limit":123,"down_speed_limt":1234,"def_visit":0,"def_visited":1}
	]`
	fmt.Fprintln(w, str)
}

//获取权限数据
func GetAccessHandler(w http.ResponseWriter, r *http.Request) {
	str := `[
		{"Name":"admin","Group1":"admin","Group2":"Boss","Rule":1},
		{"Name":"test","Group1":"guest","Group2":"Boss","Rule":0}
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

	http.Handle("/", http.FileServer(http.Dir("html")))

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}
