package qesygo

import (
	"net"
)

var UdpConn *net.UDPConn

func SocketListenUdp(Ip string, f func()) {
	if udpAddr, err := net.ResolveUDPAddr("udp", Ip); err == nil {
		if conn, err := net.ListenUDP("udp", udpAddr); err == nil {
			defer conn.Close()
			UdpConn = conn
			for {
				f()
			}
		}
	}
	Die("UDP Socket Listen Err .")
}
