/*
 * @Description: Do not edit
 * @Author: taowentao
 * @Date: 2019-01-11 13:12:59
 * @LastEditors: taowentao
 * @LastEditTime: 2019-01-11 13:30:08
 */
package tool

import "net"

type ARP struct {
	SrcIP  net.IP
	SrcMAC net.HardwareAddr
	DstIP  net.IP
	DstMAC net.HardwareAddr
}

func NewARP(data []byte) (arp *ARP) {
	arp = new(ARP)
	arp.SrcMAC = data[8:14]
	arp.SrcIP = data[14:18]
	arp.DstMAC = data[18:24]
	arp.DstIP = data[24:28]
	return
}

//是ARP
func ISARP(data []byte) (r bool) {
	if data[12] == 0x08 && data[13] == 0x06 {
		r = true
	}
	return
}

func (this *ARP) ToBin() (data []byte) {
	data = make([]byte, 28)
	data[0] = 0x00
	data[1] = 0x01
	data[2] = 0x08
	data[3] = 0x00
	data[4] = 0x06
	data[5] = 0x04
	data[6] = 0x00
	data[7] = 0x07
	copy(data[8:14], this.SrcMAC[:6])
	copy(data[14:18], this.SrcIP[:4])
	copy(data[18:24], this.DstMAC[:6])
	copy(data[24:28], this.DstIP[:4])
	return
}
