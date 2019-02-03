/*
 * @Description: 数据处理相关
 * @Author: taowentao
 * @Date: 2019-01-06 17:36:56
 * @LastEditors: taowentao
 * @LastEditTime: 2019-02-03 16:57:07
 */
package main

import (
	"client"
	"fmt"
	"net"
	"tool"
)

/**
 * @description: 处理接收到的数据,解压|解密
 * @param {type}
 * @return:
	res:处理完的数据
	ok:是否有效数据
*/
func handle_recv_data(data []byte) (res []byte, ok bool) {
	res = data
	// buff := bytes.NewReader(data)
	// var err error
	// res, _ = lzo.Decompress1X(buff, len(data), 20480)
	// if err != nil {
	// 	fmt.Println("Decompress1X ", err)
	// }
	// fmt.Println("Decompress1X ", len(data), "=>", len(res))
	ok = true
	return
}

/**
 * @description: 处理接要发送的数据,加密|压缩
 * @param {type}
 * @return:
	res:处理完的数据
	ok:是否有效数据
*/
func handle_send_data(data []byte) (res []byte, ok bool) {
	res = data
	// res = lzo.Compress1X(data)
	// fmt.Println("Compress1X ", len(data), "=>", len(res))
	ok = true
	return
}

//处理某个客户端的连接
func data_connhandler(conn net.Conn) (err error) {
	// fmt.Println(string(data))
	var clt *client.Client
	buf := make([]byte, 20480)
	for {
		var n int
		var err error
		if clt != nil {
			if clt.IsConn(1) {
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
		data, ok := handle_recv_data(buf[:n])
		if !ok {
			continue
		}
		if len(data) < 14 {
			continue
		}
		// var srcmac net.HardwareAddr //源MAC
		if clt == nil {
			// srcmac = tool.GetSrcMac(data)  //源MAC
			srcmac := tool.GetSrcMac(data) //源MAC
			// fmt.Println("srcmac:", srcmac)
			clt = client.GetClient(srcmac) //找到源客户端
			if clt == nil {
				// fmt.Println("no this src client:", srcmac)
				// client.PrintClientMap(-1)
				continue
			}
		}
		if !clt.IsConn(1) {
			fmt.Println("set data conn:", clt.Mac)
			clt.SetConn(1, conn)
		}
		clt.UpDate()

		dstmac := tool.GetDstMac(data) //目的MAC
		// fmt.Println("srcmac:", clt.Mac, "dstmac", dstmac)
		if tool.ISBroadCastMac(dstmac) { //广播
			client.BroadCast(data)
			continue
		}
		if tool.ISARP(data) {
			// arp := tool.NewARP(data)

			// pak := arp.ToBin()
		}

		dc := client.GetClient(dstmac) //找到目的客户端
		if dc == nil {
			// fmt.Println("no this client:", dstmac)
			continue
		}
		if !dc.IsConn(1) {
			// fmt.Println("this client not ready:", dstmac)
			continue
		}
		data, ok = handle_send_data(data)
		if !ok {
			continue
		}
		n, err = dc.Write(data)
		if err != nil || n < 0 {
			// fmt.Println("write err", err)
			dc.Offline()
			break
		} else {
			// fmt.Println("write ", n)
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
		fmt.Println("conn from:", conn.RemoteAddr().String())
		go data_connhandler(conn)
	}
}
