package client

import (
	"fmt"
	"net"
	"time"
)

type Client struct {
	Mac    net.HardwareAddr //vpn MAC
	IP     net.IP           //vpn IP
	conn   net.Conn         //连接
	uptime time.Time        //最后一次数据传输时间
}

func NewClient(mac net.HardwareAddr, ip net.IP, conn net.Conn) (client *Client) {
	client = &Client{
		Mac:  mac,
		IP:   ip,
		conn: conn,
	}
	return
}

//上线，添加到map
func (this *Client) Online() (r bool) {
	clientmap.Store(this.Mac.String(), this)
	fmt.Println(this.Mac, "online")
	return
}

//下线，从map删除
func (this *Client) Offline() (r bool) {
	clientmap.Delete(this.Mac.String())
	fmt.Println(this.Mac, "offline")
	return
}

//写入
func (this *Client) Write(data []byte) (n int, err error) {
	n, err = this.conn.Write(data)
	// fmt.Println("write to ", this.Mac, n, "bytes")
	if n <= 0 {
		this.Offline()
	}
	return
}

//读取
func (this *Client) Read(buf []byte) (n int, err error) {
	n, err = this.conn.Read(buf)
	if n <= 0 {
		this.Offline()
	}
	return
}

/**
 * @description: 更新时间
 * @param {type}
 * @return:
 */
func (this *Client) UpDate() {
	this.uptime = time.Now()
}
