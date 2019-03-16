package client

import (
	// "fmt"
	"net"
)

type Client struct {
	Mac  net.HardwareAddr //MAC
	Conn *net.UDPConn     //连接
	Addr *net.UDPAddr     //地址
}

func NewClient(mac net.HardwareAddr, conn *net.UDPConn, addr *net.UDPAddr) (client *Client) {
	client = &Client{
		Mac:  mac,
		Conn: conn,
		Addr: addr,
	}
	return
}

// //设置连接,1为数据，2为控制
// func (this *Client) SetConn(tp int, con net.Conn) {
// 	switch tp {
// 	case 1:
// 		this.conn = con
// 	case 2:
// 		this.connctl = con
// 	}
// }

// //是否连接,1为数据，2为控制
// func (this *Client) IsConn(tp int) (r bool) {
// 	switch tp {
// 	case 1:
// 		r = this.conn != nil
// 	case 2:
// 		r = this.connctl != nil
// 	}
// 	return
// }

//写入
func (this *Client) Write(data []byte) (n int, err error) {
	n, err = this.Conn.WriteToUDP(data, this.Addr)
	// fmt.Println("write to ", this.Mac, n, "bytes")
	return
}

//读取
func (this *Client) Read(buf []byte) (n int, err error) {
	// fmt.Println(this.conn)
	n, this.Addr, err = this.Conn.ReadFromUDP(buf)
	return
}

/**
 * @description: 更新时间
 * @param {type}
 * @return:
 */
func (this *Client) UpDate() {
	// this.uptime = time.Now()
}
