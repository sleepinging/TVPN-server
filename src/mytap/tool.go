package main

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

//是ARP
func ISARP(date []byte) (r bool) {

	return
}
