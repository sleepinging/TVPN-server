package main

import (
	"client"
	"fmt"
	"net"
	"time"
	"tool"
)

//处理接收到的数据
func data_handler(data []byte, conn *net.UDPConn, addr *net.UDPAddr) (err error) {
	// fmt.Println(string(data))
	if len(data) < 14 {
		return
	}
	// c := client.NewOnlineClient(
	// 	tool.GetSrcMac(data),
	// 	conn,
	// 	addr,
	// )
	c := client.GetOnlineClient(tool.GetSrcMac(data))
	if c == nil {
		return
	}
	c.Conn = conn
	c.Addr = addr
	c.Update = time.Now()

	// if client.GetOnlineClient(c.Mac) == nil {
	// 	if client.IsAllow_online(c.Mac) {
	// 		c.Online()
	// 		// c.Info=
	// 		fmt.Println(c.Mac.String(), "Online")
	// 		// client.PrintOnlineClientMap(-1)
	// 	} else {
	// 		fmt.Println(c.Mac.String(), "not allow online")
	// 	}
	// 	return
	// }

	dstmac := tool.GetDstMac(data) //找到目的客户端
	if tool.ISBroadCastMac(dstmac) {
		client.BroadCast(data)
	}

	dc := client.GetOnlineClient(dstmac)
	if dc == nil {
		return
	}
	n, err := dc.Write(data)
	if err != nil || n < 0 {
		dc.Offline()
		fmt.Println(c.Mac.String(), "offline")
	}

	return
}

//监听端口数据协程
func th_listen_data(port int) {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("0.0.0.0"), Port: port})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	data := make([]byte, 65535)
	for {
		n, addr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// fmt.Println("read from:", addr)
		data_handler(data[:n], conn, addr)
	}

}
