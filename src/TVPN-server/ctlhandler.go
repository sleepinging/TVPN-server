/*
 * @Description: 客户端控制相关,上下线等
 * @Author: taowentao
 * @Date: 2019-01-06 17:37:40
 * @LastEditors: taowentao
 * @LastEditTime: 2019-03-16 18:35:18
 */
package main

import (
	"fmt"
	"net"

	"client"
)

//处理接收到的数据
func ctl_handler(data []byte, conn *net.UDPConn, addr *net.UDPAddr) (err error) {
	// fmt.Println(string(data))
	//解密
	if len(data) < 1 { //长度不对
		return
	}
	switch data[0] {
	case 0x01: //登录
		err = login_handler(data[1:], conn, addr)
	case 0x02: //DHCP

	}

	return
}

//处理登录
func login_handler(data []byte, conn *net.UDPConn, addr *net.UDPAddr) (err error) {
	// fmt.Println("recv login ...")
	//MAC IP KEY
	if len(data) != 42 { //长度不对
		return
	}
	var mac net.HardwareAddr = data[0:6]
	var ip net.IP = data[6:10]
	key := data[11:43]
	//检查key
	key = key
	//检查IP和MAC
	ip = ip
	mac = mac
	//先把原来的下线
	client.Offline(mac)
	if client.Allow_online(mac, ip) {
		fmt.Println(mac.String(), "allow online")
	}
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
