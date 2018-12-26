package client

import (
	"fmt"
	"net"
	"sync"
)

//在线客户端map,mac[*client]
var clientmap = sync.Map{}

type Client struct {
	Mac  net.HardwareAddr //MAC
	Conn *net.UDPConn     //连接
	Addr *net.UDPAddr     //地址
}

func init() {

}

//上线，添加到map
func (this *Client) Online() (r bool) {
	clientmap.Store(this.Mac, this)
	fmt.Println(this.Mac, "online")
	return
}

//下线，从map删除
func (this *Client) Offline() (r bool) {
	clientmap.Delete(this.Mac)
	return
}

//打印map,调试时用,最大条数-1无限制
func PrintClientMap(max int) {
	clientmap.Range(func(k, v interface{}) bool {
		max--
		if max == -1 {
			return false
		}
		fmt.Printf("mac:%+v,client:%+v\n", k.(net.HardwareAddr), v.(*Client))
		return true
	})
}

//获取客户端
func GetClient(mac net.HardwareAddr) (c *Client) {
	ic, ok := clientmap.Load(mac)
	if !ok {
		return
	}
	c = ic.(*Client)
	return
}
