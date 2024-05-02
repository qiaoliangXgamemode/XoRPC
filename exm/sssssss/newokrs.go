package sssssss

import(
	"net"
	"log"
)


// Create TCP Conn Dail
func CreateTCPConn(Addres string) (*net.TCPConn) {
	if Addres == "" { return nil }
	tcpAddr, _ := net.ResolveTCPAddr("tcp", Addres)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Printf("[Info][Control][CreateTCPConn][ERROR] Dial TCP error: %s", err)
	}
	return conn
}
// Create TCP net.Listen Listen loacl Addres
func CreateTCPListen(Addres string) (net.Listener) {
	if Addres == "" { return nil }
	l, err := net.Listen("tcp", Addres)
	if err != nil {
		log.Printf("[Info][Control][CreateTCPListen][ERROR] Listen TCP error: %s", err)
	}
	return l
}

// Create UDP net.UDPConn
func CreateUDPConn(Addres string) (*net.UDPConn, error) {
	if Addres == "" { return nil, nil }
	RemoteAddres,_ := net.ResolveUDPAddr("udp", Addres)
	conn, err := net.DialUDP("udp", nil, RemoteAddres)
	if err != nil {
		log.Printf("[Info][Control][CreateUDPConn][ERROR] Dail UDP error: %s", err)
		return nil, err
	}
	return conn, nil
}

// Create UDP net.UDPAddr
func CreateUDPListen(Addres string) (*net.UDPConn, error) {
	if Addres == "" { return nil, nil }
	srcAddr,_ := net.ResolveUDPAddr("udp", Addres)
	conn, err := net.ListenUDP("udp", srcAddr)
	if err != nil {
		log.Printf("[Info][Control][CreateUDPListen][ERROR] Listen UDP error: %s", err)
		return nil, err
	}
	return conn, nil
}

// Create KCP kcp.seeosin
// func CreateKCPConn(Addres string) {
// 	if Addres == "" { return nil, nil }
// }
// // Create KCP kcp.seeosin
// func CreateKCPListen(Addres string) {
// 	if Addres == "" { return nil, nil }
// }

// QUIC TODO
// func CreateQuicConn(Addre string) {
// 	if Addres == nil {
		
// 	}
// }

// func CreateQuicListen(Addre string) {
// 	if Addres == nil {
		
// 	}
// }