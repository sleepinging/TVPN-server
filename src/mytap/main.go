package main

import (
	"fmt"
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

	fmt.Println("data port", dataport)
	th_listen_data(dataport)
}
