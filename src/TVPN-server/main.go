package main

import (
	"client"
	"fmt"
	"time"
	"view"
)

func main() {
	ctlport := 6544
	dataport := 6543
	httpport := 6580

	fmt.Println("http port", httpport)
	go view.StartHTTPServer(httpport)

	fmt.Println("control port", ctlport)
	go th_listen_ctl(ctlport)

	//10分钟检查一次，5小时不不活跃下线
	go client.Clean_Client(time.Second*60*10, time.Second*60*60*5)

	fmt.Println("data port", dataport)
	th_listen_data(dataport)
}
