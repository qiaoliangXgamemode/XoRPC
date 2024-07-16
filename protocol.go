package XoRPC

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"log/slog"
	"net"

	"github.com/quic-go/quic-go"
)

var QUICVersionContextKey []quic.VersionNumber

// (*Listener).Addr
// Create TCP net.Listen Listen loacl Addres
func CreateTCPListen(Addres string) (lis *net.Listener, e error) {
	listen, err := net.Listen("tcp", Addres)
	if err != nil {
		slog.Error("[Info][Control][CreateTCPListen][ERROR] CreateUDPListenListen TCP error: %s", err)
	}
	return &listen, nil
}

// (*Listener).Addr
// Create UDP net.UDPAddr
func CreateUDPListen(Addres string) (listen *net.UDPConn, e error) {
	if Addr, e := net.ResolveUDPAddr("udp", Addres); e == nil {
		if listen, e := net.ListenUDP("udp", Addr); e == nil {
			return listen, nil
		} else {
			slog.Error("[Info][Control][CreateUDPListen][ERROR] Listen UDP error: %s", e)
			return nil, e
		}
	} else {
		slog.Error("[Info][Control][CreateUDPListen][ERROR] Listen UDP error: %s", e)
		return nil, e
	}
	// return nil, nil
}

// (*Listener).Addr
// Create KCP kcp.seeosin
// func CreateKCPListen(Addres string) (lis *kcp.Listener, e error) {
// 	if kcps, err := kcp.ListenWithOptions(Addres, nil, 10, 3); err == nil {
// 		return kcps, nil
// 	} else {
// 		log.Fatal("[Info][Control][CreateKCPListen][ERROR] Listen KCP error: %s", err)
// 	}
// 	return nil, nil
// }

// (*Conn).Addr
// Create TCP Conn Dail
func CreateTCPConn(Addres string) (dail net.Conn, e error) {
	conn, err := net.Dial("tcp", Addres)
	if err != nil {
		log.Printf("[Info][Control][CreateTCPConn][ERROR] Dial TCP error: %s", err)
	}
	return conn, nil
}

// (*Conn).Addr
// Create UDP net.UDPConn
func CreateUDPConn(Addres string) (dail net.Conn, e error) {
	if conn, e := net.Dial("udp", Addres); e == nil {
		return conn, nil
	} else {
		slog.Error("[Control][CreateUDPConn] Dail UDP error: %s", e)
		return nil, e
	}
}

// (*Conn).Addr
// Create KCP kcp.seeosin
// func CreateKCPConn(Addres string) (lis *kcp.UDPSession, e error) {
// 	if conn, err := kcp.DialWithOptions(Addres, nil, 10, 3); err == nil {
// 		return conn, nil
// 	} else {
// 		log.Fatal("[Info][Control][CreateKCPConn][ERROR] Listen KCP error: %s", err)
// 	}
// 	return nil, nil
// }

// QUIC TODO
func CreateQuicConn(Addres string) (c quic.Connection, e error) {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"XoRPC"},
	}
	conn, e := quic.DialAddr(context.Background(), Addres, tlsConf, nil)
	// conn, e := quic.DialContext(context.Background(), conn, remoteAddr, conn.LocalAddr().String(), tlsConf,
	// &quic.Config{Versions: quicVersion, MaxIdleTimeout: idleTimeout, DisablePathMTUDiscovery: true})
	if e != nil {
		// panic(e)
		slog.Error("[Control][CreateQuicConn][ERROR] Dail QUIC error: %s", e)
		return nil, e
	}
	slog.Info("[Control][CreateQuicConn] Dail QUIC ok.")
	return conn, nil
}

func CreateQuicListen(Addres string, port int) (q *quic.Listener, e error) {
	// addr := "localhost:8080"
	// &quic.Config{Versions: QUICVersionContextKey, MaxIdleTimeout: idleTimeout, DisablePathMTUDiscovery: true}
	fmt.Sprintf("%s:%d", Addres, port)
	l, e := quic.ListenAddr(Addres, generateTLSConfig(), nil)
	if e != nil {
		// panic(e)
		slog.Error("[Control][CreateQuicConn][ERROR] Dail UDP error: %s", e)
		return nil, e
	}
	slog.Info("[Control][CreateQuicConn] Dail UDP ok.")
	return l, nil
}

func generateTLSConfig() *tls.Config {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		panic(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	config.NextProtos = append(config.NextProtos, quic.Version2.String())

	return config
}
