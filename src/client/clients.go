package client

import (
	"fmt"
	"net"
	"sync"
	"time"
	"tool"
)

//在线客户端map,mac[*client]
var clientmap = sync.Map{}

//获取全部用户,最大条数
func GetAllOnlineClient(max int) (clients []*OnlineClient) {
	clients = make([]*OnlineClient, 0, max)
	i := 0
	clientmap.Range(func(k, v interface{}) bool {
		if i >= max {
			return false
		}
		clients = append(clients, v.(*OnlineClient))
		i++
		return true
	})
	return
}

//打印map,调试时用,最大条数-1无限制
func PrintOnlineClientMap(max int) {
	clientmap.Range(func(k, v interface{}) bool {
		max--
		if max == -1 {
			return false
		}
		fmt.Printf("mac:%+v,client:%+v\n", tool.I642MAC(k.(int64)), v.(*OnlineClient))
		return true
	})
}

//上线，添加到map
func (this *OnlineClient) Online() (r bool) {
	clientmap.Store(tool.MAC2I64(this.Mac), this)
	// fmt.Println(this.Mac, "online")
	return true
}

//下线，从map删除
func (this *OnlineClient) Offline() (r bool) {
	clientmap.Delete(tool.MAC2I64(this.Mac))
	// fmt.Println(this.Mac, "offline")
	return true
}

//获取客户端
func GetOnlineClient(mac net.HardwareAddr) (c *OnlineClient) {
	ic, ok := clientmap.Load(tool.MAC2I64(mac))
	if !ok {
		return
	}
	c = ic.(*OnlineClient)
	return
}

//下线，从map删除
func Offline(mac net.HardwareAddr) (r bool) {
	clientmap.Delete(tool.MAC2I64(mac))
	return true
}

//定期清除不活跃客户端
func Clean_Client(dur, timeout time.Duration) {
	for {
		clientmap.Range(func(k, v interface{}) bool {
			c := v.(*OnlineClient)
			if c.Update.Add(timeout).Before(time.Now()) {
				clientmap.Delete(k)
			}
			return true
		})
		time.Sleep(dur)
	}
}

//广播
func BroadCast(data []byte) {
	// fmt.Println("broad cast")
	clientmap.Range(func(k, v interface{}) bool {
		c := v.(*OnlineClient)
		// if c.IsConn(1) { //已经建立数据连接
		c.Write(data)
		// }
		return true
	})
}
