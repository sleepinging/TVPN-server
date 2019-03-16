/*
 * @Description: 网络相关的函数
 * @Author: taowentao
 * @Date: 2018-12-26 19:52:57
 * @LastEditors: taowentao
 * @LastEditTime: 2019-03-16 18:29:39
 */
package tool

import (
	"net"
)

//从数据包获取目的MAC
func GetDstMac(data []byte) (mac net.HardwareAddr) {
	mac = []byte{data[0], data[1], data[2], data[3], data[4], data[5]}
	return
}

//从数据包获取源MAC
func GetSrcMac(data []byte) (mac net.HardwareAddr) {
	mac = []byte{data[6], data[7], data[8], data[9], data[10], data[11]}
	return
}

//是广播MAC
func ISBroadCastMac(mac net.HardwareAddr) (r bool) {
	for i := 0; i < 6; i++ {
		if mac[i] != 0xff {
			return
		}
	}
	r = true
	return
}

//把MAC转int64
func MAC2I64(mac net.HardwareAddr) (v int64) {
	for i := 0; i < 5; i++ {
		v |= int64(mac[i])
		v <<= 8
	}
	v |= int64(mac[5])
	return
}

//把int64转MAC
func I642MAC(v int64) (mac net.HardwareAddr) {
	mac = make([]byte, 6)
	mac[5] = byte(v & 0x00000000000000ff)
	mac[4] = byte((v >> 8) & 0x00000000000000ff)
	mac[3] = byte((v >> 16) & 0x00000000000000ff)
	mac[2] = byte((v >> 24) & 0x00000000000000ff)
	mac[1] = byte((v >> 32) & 0x00000000000000ff)
	mac[0] = byte((v >> 40) & 0x00000000000000ff)
	return
}
