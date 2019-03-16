package client

import (
	"fmt"
	"net"
	"sync"
)

//在线客户端map,mac[*client]
var clientmap = sync.Map{}

//打印map,调试时用,最大条数-1无限制
func PrintClientMap(max int) {
	clientmap.Range(func(k, v interface{}) bool {
		max--
		if max == -1 {
			return false
		}
		fmt.Printf("mac:%+v,client:%+v\n", k.(string), v.(*Client))
		return true
	})
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

//获取客户端
func GetClient(mac net.HardwareAddr) (c *Client) {
	ic, ok := clientmap.Load(mac.String())
	if !ok {
		return
	}
	c = ic.(*Client)
	return
}

//广播
func BroadCast(data []byte) {
	// fmt.Println("broad cast")
	clientmap.Range(func(k, v interface{}) bool {
		c := v.(*Client)
		// if c.IsConn(1) { //已经建立数据连接
		c.Write(data)
		// }
		return true
	})
}
