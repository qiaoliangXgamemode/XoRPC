package XoRPC

import (
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/quic-go/quic-go"
)

type protocolQUIC struct {
	l        quic.Listener
	writeMtx *sync.Mutex
	quic.Stream
	quic.Connection
}

type protocolTCP struct {
	l        net.Listener
	writeMtx *sync.Mutex
}

type protocolUDP struct {
	l        *net.UDPConn
	writeMtx *sync.Mutex
}

type protocolWebsocket struct {
	l               *websocket.Conn
	w               io.Writer
	r               io.Reader
	ReadBufferSize  int
	WriteBufferSize int
	// writeMtx *sync.Mutex
}

func ListenerQUIC(addr string, port int) *protocolQUIC {
	lis, _ := CreateQuicListen(addr, port)
	n := new(protocolQUIC)
	n.l = *lis
	return n
}

func ListenerTCP(addr string, port int) *protocolTCP {
	lis, _ := CreateTCPListen(fmt.Sprintf("%s:%d", addr, port))
	n := new(protocolTCP)
	n.l = lis
	return n
}

func ListenerUDP(addr string, port int) *protocolUDP {
	lis, _ := CreateUDPListen(fmt.Sprintf("%s:%d", addr, port))
	n := new(protocolUDP)
	n.l = lis
	return n
}

func ListenWebsocket(addr string, port int, readBuff int, writeBuff int) *protocolWebsocket {
	lis, _ := CreateWebsocketListen(addr, port, readBuff, writeBuff)
	n := new(protocolWebsocket)
	n.l = lis
	n.ReadBufferSize = readBuff
	n.WriteBufferSize = writeBuff
	return n
}
