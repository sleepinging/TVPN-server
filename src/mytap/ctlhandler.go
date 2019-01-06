/*
 * @Description: 客户端控制相关,上下线等
 * @Author: taowentao
 * @Date: 2019-01-06 17:37:40
 * @LastEditors: taowentao
 * @LastEditTime: 2019-01-06 17:43:20
 */
package main

import (
	"fmt"
	"net"
)

//处理某个客户端的连接
func ctl_connhandler(conn net.Conn) (err error) {
	return
}

//监听控制端口协程
func th_listen_ctl(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		fmt.Printf("listen fail, err: %v\n", err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept fail, err: %v\n", err)
			continue
		}
		// fmt.Println("conn from:", conn)
		go ctl_connhandler(conn)
	}
}
