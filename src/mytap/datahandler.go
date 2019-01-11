/*
 * @Description: 数据处理相关
 * @Author: taowentao
 * @Date: 2019-01-06 17:36:56
 * @LastEditors: taowentao
 * @LastEditTime: 2019-01-11 13:08:34
 */
package main

import (
	"client"
	"fmt"
	"net"
	"tool"
)

/**
 * @description: 处理接收到的数据,解压|解密,加密|压缩
 * @param {type}
 * @return:
	res:处理完的数据
	ok:是否有效数据
*/
func handle_recv_data(data []byte) (res []byte, ok bool) {
	res = data
	ok = true
	return
}

//处理某个客户端的连接
func data_connhandler(conn net.Conn) (err error) {
	// fmt.Println(string(data))
	var clt *client.Client = nil
	buf := make([]byte, 2048)
	for {
		var n int
		var err error
		if clt != nil {
			n, err = clt.Read(buf)
		} else {
			n, err = conn.Read(buf)
		}
		if err != nil {
			fmt.Println("read", err)
			break
		}
		data, ok := handle_recv_data(buf[:n])
		if !ok {
			continue
		}
		if len(data) < 14 {
			continue
		}

		if clt == nil {
			srcmac := tool.GetSrcMac(data) //源MAC
			clt = client.GetClient(srcmac) //找到源客户端
			if clt == nil {
				clt = client.NewClient(srcmac, nil, conn)
				clt.Online()
				// client.PrintClientMap(-1)
			}
		}
		clt.UpDate()

		dstmac := tool.GetDstMac(data) //目的MAC
		// fmt.Println("srcmac:", srcmac, "dstmac", dstmac)
		if tool.ISBroadCastMac(dstmac) { //ARP
			client.BroadCast(data)
			continue
		}

		dc := client.GetClient(dstmac) //找到目的客户端
		if dc == nil {
			// fmt.Println("no this client:", dstmac)
			continue
		}
		n, err = dc.Write(data)
		if err != nil {
			fmt.Println("write", err)
			break
		}
	}

	return
}

//监听数据端口协程
func th_listen_data(port int) {
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
		go data_connhandler(conn)
	}
}
