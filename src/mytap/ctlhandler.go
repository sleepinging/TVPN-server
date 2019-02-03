/*
 * @Description: 客户端控制相关,上下线等
 * @Author: taowentao
 * @Date: 2019-01-06 17:37:40
 * @LastEditors: taowentao
 * @LastEditTime: 2019-02-03 16:08:05
 */
package main

import (
	"client"
	"fmt"
	"net"
)

//处理某个客户端的连接
func ctl_connhandler(conn net.Conn) (err error) {
	var clt *client.Client
	buf := make([]byte, 2048)
	for {
		var n int
		var err error
		if clt != nil {
			if clt.IsConn(2) {
				n, err = clt.Read(buf)
				if n <= 0 || err != nil {
					clt.Offline()
				}
			}
		} else {
			n, err = conn.Read(buf)
		}
		if err != nil {
			// fmt.Println("read err:", err)
			break
		}
		data := buf[:n]
		// data, ok := handle_recv_data(buf[:n])
		// if !ok {
		// 	continue
		// }
		// if len(data) < 14 {
		// 	continue
		// }
		if len(data) < 1 {
			fmt.Println("ctl data too short")
			continue
		}
		switch data[0] {
		case 1: //心跳
			handle_heart_beat(clt, data[1:], conn)
		}
	}

	return
}

//处理心跳包
func handle_heart_beat(clt *client.Client, macdata []byte, conn net.Conn) {
	// fmt.Println("recv heart beat")
	if clt == nil {
		var srcmac net.HardwareAddr //源MAC
		if len(macdata) < 6 {
			fmt.Println("macdata too short")
			return
		}
		srcmac = macdata[0:6]
		clt = client.GetClient(srcmac) //找到源客户端
		if clt == nil {
			clt = client.NewClient(srcmac, nil)
			clt.Online()
			// client.PrintClientMap(-1)
		}
		if !clt.IsConn(2) {
			fmt.Println("set ctl conn:", clt.Mac)
			clt.SetConn(2, conn)
		}
	}
	clt.UpDate()
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
		fmt.Println("conn from:", conn.RemoteAddr().String())
		go ctl_connhandler(conn)
	}
}
