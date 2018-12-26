package main

import (
	"client"
	"fmt"
	"net"
)

//处理接收到的数据
func datahandler(data []byte, conn *net.UDPConn, addr *net.UDPAddr) (err error) {
	fmt.Println(string(data))
	if len(data) < 14 {
		return
	}
	c := client.Client{
		Mac:  GetSrcMac(data),
		Conn: conn,
		Addr: addr,
	}
	if client.GetClient(c.Mac) == nil {
		c.Online()
	}
	dstmac := GetDstMac(data) //找到目的客户端
	dc := client.GetClient(dstmac)
	if dc == nil {
		return
	}
	// dc.Addr.

	return
}

//监听端口数据协程
func th_listen() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("0.0.0.0"), Port: 6543})
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
		fmt.Println("read from:", addr)
		datahandler(data[:n], conn, addr)
	}

}

func main() {
	fmt.Println("开始...")
	th_listen()
}
