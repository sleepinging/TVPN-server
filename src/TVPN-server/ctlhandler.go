/*
 * @Description: 客户端控制相关,上下线等
 * @Author: taowentao
 * @Date: 2019-01-06 17:37:40
 * @LastEditors: taowentao
 * @LastEditTime: 2019-05-19 16:21:29
 */
package main

import (
	"controller"
	"encoding/binary"
	"fmt"
	"net"
	"time"
	"tool"

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
	if len(data) < 15 { //长度不对
		return
	}
	idx := 0
	var mac net.HardwareAddr = data[idx : idx+6]
	idx += 6
	var ip net.IP = data[idx : idx+4]
	idx += 4
	l := binary.BigEndian.Uint16(data[idx:])
	idx += 2
	if idx+int(l) > len(data) {
		return
	}
	name := string(data[idx : idx+int(l)])
	idx += int(l)
	l = binary.BigEndian.Uint16(data[idx:])
	if idx+int(l)+2 > len(data) {
		return
	}
	idx += 2
	pwd := string(data[idx : idx+int(l)])
	idx += int(l)
	// fmt.Println(string(name), string(pwd))
	//检查用户名和密码
	if !controller.CheckLogin(name, pwd) {
		conn.WriteToUDP([]byte{0x00, 0x02}, addr)
		return
	}
	// key := data[11:43]
	// //检查key
	// key = key
	//TODO:检查IP和MAC
	ip = ip
	mac = mac
	//先把原来的下线
	// client.Offline(mac)
	if client.Allow_online(mac, ip) {
		c := client.OnlineClient{
			Mac:     tool.CopyMAC(mac),
			ConnCTL: conn,
			AddrCTL: addr,
			Info: &client.Client{
				Name: name,
			},
			IP:         tool.CopyIP(ip),
			Onlinetime: time.Now(),
		}
		fmt.Println(c.ToJson(), "online")
		c.Online()
		conn.WriteToUDP([]byte{0x00, 0x01}, addr)
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
