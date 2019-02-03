package client

import (
	"fmt"
	"net"
	"sync"
	"time"
)

//在线客户端map,macstring[*client]
var clientmap = sync.Map{}
var timeout_offline = time.Minute

func init() {
	go th_clean_client(time.Second * 30)
}

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
		if c.IsConn(1) { //已经建立数据连接
			c.Write(data)
		}
		return true
	})
}

/**
 * @description: 定时清理长期没数据的设备
 * @param {type}
 * @return:
 */
func th_clean_client(dur time.Duration) {
	for {
		clientmap.Range(func(k, v interface{}) bool {
			c := v.(*Client)
			if c.uptime.Add(timeout_offline).Before(time.Now()) {
				fmt.Println(c.Mac, timeout_offline, "no update,offline")
				c.Offline()
			}
			return true
		})
		time.Sleep(dur)
	}
}
