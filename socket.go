package QesyGo

import (
	"net"
)

var UdpConn *net.UDPConn

func SocketListenUdp(Ip string, Rec func()) {
	if udpAddr, err := net.ResolveUDPAddr("udp", Ip); err == nil {
		if conn, err := net.ListenUDP("udp", udpAddr); err == nil {
			defer conn.Close()
			UdpConn = conn
			for {
				Rec()
			}
		}
	}
	Die("UDP Socket Listen Err .")
}
