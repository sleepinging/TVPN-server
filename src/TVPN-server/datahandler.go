package main

import (
	"client"
	"fmt"
	"net"
	"tool"
)

//处理接收到的数据
func datahandler(data []byte, conn *net.UDPConn, addr *net.UDPAddr) (err error) {
	// fmt.Println(string(data))
	if len(data) < 14 {
		return
	}
	c := client.NewClient(
		tool.GetSrcMac(data),
		conn,
		addr,
	)

	if client.GetClient(c.Mac) == nil {
		c.Online()
		fmt.Println(c.Mac.String(), "online")
	}

	dstmac := tool.GetDstMac(data) //找到目的客户端
	if tool.ISBroadCastMac(dstmac) {
		client.BroadCast(data)
	}

	dc := client.GetClient(dstmac)
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
	data := make([]byte, 2048)
	for {
		n, addr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// fmt.Println("read from:", addr)
		datahandler(data[:n], conn, addr)
	}

}
