/*
 * @Description: 客户端控制相关,上下线等
 * @Author: taowentao
 * @Date: 2019-01-06 17:37:40
 * @LastEditors: taowentao
 * @LastEditTime: 2019-03-16 16:16:51
 */
package main

import (
	"fmt"
	"net"
)

//处理接收到的数据
func ctl_handler(data []byte, conn *net.UDPConn, addr *net.UDPAddr) (err error) {
	// fmt.Println(string(data))

	return
}

//监听控制端口协程
func th_listen_ctl(port int) {
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
		ctl_handler(data[:n], conn, addr)
	}
}
