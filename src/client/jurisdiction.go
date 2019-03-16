package client

import (
	"net"
	"sync"
)

//客户端权限

//允许上线的客户端 mac-bool
var allow_clientmap = sync.Map{}

//允许上线
func Allow_online(mac net.HardwareAddr, ip net.IP) (r bool) {
	allow_clientmap.Store(mac.String(), true)
	r = true
	return
}

//检查是否允许上线
func IsAllow_online(mac net.HardwareAddr) (r bool) {
	a, ok := allow_clientmap.Load(mac.String())
	if ok {
		r = a.(bool)
	}

	return
}
