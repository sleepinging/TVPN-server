package client

import (
	// "fmt"
	"fmt"
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

	IP net.IP //使用的IP

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

//把一个用户转成json
func (c *OnlineClient) ToJson() (str string) {
	str = fmt.Sprintf(`{"Name":"%s","IP":"%s","MAC":"%s","OnlineTime":"%s"}`, c.Info.Name, c.IP.String(), c.Mac.String(), c.Onlinetime.Format("2006-01-02 15:04:05"))
	return
}

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
