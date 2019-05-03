package client

import (
	// "fmt"
	"net"
	"time"
)

type Client struct {
	Name string //用户名
}

type OnlineClient struct {
	Mac  net.HardwareAddr //MAC
	Conn *net.UDPConn     //数据连接
	Addr *net.UDPAddr     //数据地址

	ConnCTL *net.UDPConn //控制连接
	AddrCTL *net.UDPAddr //控制地址

	Onlinetime time.Time //上线时间
	Update     time.Time //活跃时间

	Info *Client //客户端信息
}

func NewOnlineClient(mac net.HardwareAddr, conn *net.UDPConn, addr *net.UDPAddr) (client *OnlineClient) {
	client = &OnlineClient{
		Mac:  mac,
		Conn: conn,
		Addr: addr,
	}
	return
}

// //设置连接,1为数据，2为控制
// func (this *OnlineClient) SetConn(tp int, con net.Conn) {
// 	switch tp {
// 	case 1:
// 		this.conn = con
// 	case 2:
// 		this.connctl = con
// 	}
// }

// //是否连接,1为数据，2为控制
// func (this *OnlineClient) IsConn(tp int) (r bool) {
// 	switch tp {
// 	case 1:
// 		r = this.conn != nil
// 	case 2:
// 		r = this.connctl != nil
// 	}
// 	return
// }

//写入
func (this *OnlineClient) Write(data []byte) (n int, err error) {
	n, err = this.Conn.WriteToUDP(data, this.Addr)
	// fmt.Println("write to ", this.Mac, n, "bytes")
	return
}

//读取
func (this *OnlineClient) Read(buf []byte) (n int, err error) {
	// fmt.Println(this.conn)
	n, this.Addr, err = this.Conn.ReadFromUDP(buf)
	return
}

/**
 * @description: 更新时间
 * @param {type}
 * @return:
 */
func (this *OnlineClient) UpDate() {
	// this.uptime = time.Now()
}
