package XoRPC

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"kcp"
	"log"
	"log/slog"
	"math/big"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/quic-go/quic-go"
)

var QUICVersionContextKey []quic.VersionNumber

// ----------------------------------------------------------------
// |                        TCP                                 |
// ----------------------------------------------------------------
// (*Listener).Addr
// Create TCP net.Listen Listen loacl Addres
func CreateTCPListen(Addres string) (lis net.Listener, e error) {
	listen, err := net.Listen("tcp", Addres)
	if err != nil {
		slog.Error("[Info][Control][CreateTCPListen][ERROR] CreateUDPListenListen TCP error: %s", err)
	}
	return listen, nil
}

// (*Conn).Addr
// Create TCP Conn Dail
func CreateTCPConn(Addres string) (dail net.Conn, e error) {
	conn, err := net.Dial("tcp", Addres)
	if err != nil {
		log.Printf("[Info][Control][CreateTCPConn][ERROR] Dial TCP error: %s", err)
	}
	return conn, nil
}

// ----------------------------------------------------------------
// |                        UDP                                 |
// ----------------------------------------------------------------
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

// ----------------------------------------------------------------
// |                        KCP                                 |
// ----------------------------------------------------------------
// (*Conn).Addr
// Create KCP kcp.seeosin
func CreateKCPConn(Addres string) (conn *kcp.UDPSession, e error) {
	if c, err := kcp.DialWithOptions(Addres, nil, 10, 3); err == nil {
		return c, nil
	} else {
		log.Fatal("[Info][Control][CreateKCPConn][ERROR] Listen KCP error: %s", err)
	}
	return nil, nil
}

// (*Listener).Addr
// Create KCP kcp.seeosin
func CreateKCPListen(Addres string) (lis *kcp.Listener, e error) {
	if l, err := kcp.ListenWithOptions(Addres, nil, 10, 3); err == nil {
		return l, nil
	} else {
		log.Fatal("[Info][Control][CreateKCPListen][ERROR] Listen KCP error: %s", err)
	}
	return nil, nil
}

// ----------------------------------------------------------------
// |                         QUIC                                 |
// ----------------------------------------------------------------
// (QUIC) Create quic connection and listen.
func CreateQuicConn(Addres string, port int) (c quic.Connection, e error) {
	switch IsIPV64(Addres) {
	case 4:
		Addres = fmt.Sprintf("%s:%d", Addres, port)
	case 6:
		Addres = fmt.Sprintf("[%s]:%d", Addres, port)
	}
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"XoRPC"},
	}
	conn, e := quic.DialAddr(context.Background(), Addres, tlsConf, nil)
	// conn, e := quic.DialContext(context.Background(), conn, remoteAddr, conn.LocalAddr().String(), tlsConf,
	// &quic.Config{Versions: quicVersion, MaxIdleTimeout: idleTimeout, DisablePathMTUDiscovery: true})
	if e != nil {
		// panic(e)
		slog.Error("ERROR message", "[Control][CreateQuicConn][ERROR] Dail QUIC error", e)
		return nil, e
	}
	slog.Debug("Debug message", "Dail QUIC", "ok")
	return conn, nil
}

func CreateQuicListen(Addres string, port int) (q *quic.Listener, e error) {
	// addr := "localhost:8080"
	// &quic.Config{Versions: QUICVersionContextKey, MaxIdleTimeout: idleTimeout, DisablePathMTUDiscovery: true}
	Addres = fmt.Sprintf("%s:%d", Addres, port)
	l, e := quic.ListenAddr(Addres, generateTLSConfig(), nil)
	if e != nil {
		// panic(e)
		slog.Error("ERROR message", "Listen QUIC error", e)
		return nil, e
	}
	slog.Debug("Debug message", "Listen QUIC", "ok")
	return l, nil
}

// ----------------------------------------------------------------
// |                       WEBSOCKET                              |
// ----------------------------------------------------------------
// Listener interface and connection interface for the websocket
func CreateWebsocketListen(Addres string, port int, ReadBuff int, WriteBuff int) (c *websocket.Conn, e error) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  ReadBuff,
		WriteBufferSize: WriteBuff,
	}
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return conn, err
		}
	})
}
func CreateWebSocketConnect(Addres string, port int) (c *websocket.Conn) {
	var ws *websocket.Conn
	wsDst := fmt.Sprintf("ws://%s:%d", Addres, port)

	if conn, _, e := websocket.DefaultDialer.Dial(wsDst, nil); e != nil {
		return conn
	}
}
func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	// fmt.Printf("%s", keyPEM)
	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	// Certificates: []tls.Certificate{tlsCert},
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"XoRPC"},
	}
}

// func generateTLSConfig() *tls.Config {
// 	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
// 	if err != nil {
// 		panic(err)
// 	}

// 	config := &tls.Config{Certificates: []tls.Certificate{cert}}
// 	config.NextProtos = append(config.NextProtos, quic.Version2.String())

// 	return config
// }
