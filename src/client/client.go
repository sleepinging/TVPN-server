package client

import (
	"fmt"
	"net"
	"time"
)

type Client struct {
	Mac     net.HardwareAddr //vpn MAC
	IP      net.IP           //vpn IP
	conn    net.Conn         //数据连接
	connctl net.Conn         //控制连接
	uptime  time.Time        //最后一次数据传输时间
}

func NewClient(mac net.HardwareAddr, ip net.IP) (client *Client) {
	client = &Client{
		Mac: mac,
		IP:  ip,
	}
	return
}

//设置连接,1为数据，2为控制
func (this *Client) SetConn(tp int, con net.Conn) {
	switch tp {
	case 1:
		this.conn = con
	case 2:
		this.connctl = con
	}
}

//是否连接,1为数据，2为控制
func (this *Client) IsConn(tp int) (r bool) {
	switch tp {
	case 1:
		r = this.conn != nil
	case 2:
		r = this.connctl != nil
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
	return
}

//读取
func (this *Client) Read(buf []byte) (n int, err error) {
	// fmt.Println(this.conn)
	n, err = this.conn.Read(buf)
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
