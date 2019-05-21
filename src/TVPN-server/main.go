package main

import (
	"client"
	"config"
	"fmt"
	"time"
	"view"

	"dao"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	err := dao.InitDB()
	if err != nil {
		panic(err)
	}
}

func main() {
	ctlport := 6544
	dataport := 6543
	httpport := 6580

	config.InitConfig()

	fmt.Println("http port", httpport)
	go view.StartHTTPServer(httpport)

	fmt.Println("control port", ctlport)
	go th_listen_ctl(ctlport)

	//10分钟检查一次，5小时不不活跃下线
	go client.Clean_Client(time.Second*60*10, time.Second*60*60*5)
	// go client.Clean_Client(time.Second*10, time.Second*60)

	fmt.Println("data port", dataport)
	th_listen_data(dataport)
}
