package nat

import (
	"log"
	"net"

	"github.com/xtaci/kcp-go"
)

// (*Listener).Addr
// Create TCP net.Listen Listen loacl Addres
func CreateTCPListen(Addres string) (lis *net.Listener, e error) {
	if Addres == "" {
		return nil, nil
	}
	listen, err := net.Listen("tcp", Addres)
	if err != nil {
		log.Printf("[Info][Control][CreateTCPListen][ERROR] CreateUDPListenListen TCP error: %s", err)
	}
	return &listen, nil
}

// (*Listener).Addr
// Create UDP net.UDPAddr
func CreateUDPListen(Addres string) (lis *net.UDPConn, e error) {
	if Addres == "" {
		return nil, nil
	}
	srcAddr, _ := net.ResolveUDPAddr("udp", Addres)
	listen, err := net.ListenUDP("udp", srcAddr)
	if err != nil {
		log.Printf("[Info][Control][CreateUDPListen][ERROR] Listen UDP error: %s", err)
		return nil, err
	}
	return listen, nil
}

// (*Listener).Addr
// Create KCP kcp.seeosin
func CreateKCPListen(Addres string) (lis *kcp.Listener, e error) {
	if kcps, err := kcp.ListenWithOptions(Addres, nil, 10, 3); err == nil {
		return kcps, nil
	} else {
		log.Fatal("[Info][Control][CreateKCPListen][ERROR] Listen KCP error: %s", err)
	}
	return nil, nil
}

// (*Conn).Addr
// Create TCP Conn Dail
func CreateTCPConn(Addres string) (dail net.Conn, e error) {
	if Addres == "" {
		return nil, nil
	}
	conn, err := net.Dial("tcp", Addres)
	if err != nil {
		log.Printf("[Info][Control][CreateTCPConn][ERROR] Dial TCP error: %s", err)
	}
	return conn, nil
}

// (*Conn).Addr
// Create UDP net.UDPConn
func CreateUDPConn(Addres string) (dail net.Conn, e error) {
	if Addres == "" {
		return nil, nil
	}
	conn, err := net.Dial("udp", Addres)
	if err != nil {
		log.Printf("[Info][Control][CreateUDPConn][ERROR] Dail UDP error: %s", err)
		return nil, err
	}
	return conn, nil
}

// (*Conn).Addr
// Create KCP kcp.seeosin
func CreateKCPConn(Addres string) (lis *kcp.UDPSession, e error) {
	if conn, err := kcp.DialWithOptions(Addres, nil, 10, 3); err == nil {
		return conn, nil
	} else {
		log.Fatal("[Info][Control][CreateKCPConn][ERROR] Listen KCP error: %s", err)
	}
	return nil, nil
}

// QUIC TODO
// func CreateQuicConn(Addre string) {
// 	if Addres == nil {

// 	}
// }

// func CreateQuicListen(Addre string) {
// 	if Addres == nil {

// 	}
// }
