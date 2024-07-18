package XoRPC

import (
	"sync"

	"github.com/quic-go/quic-go"
)

type typeQUIC struct {
	l        quic.Listener
	writeMtx *sync.Mutex
	quic.Stream
	quic.Connection
}

func ListenerQUIC(addr string, port int) *typeQUIC {
	lis, _ := CreateQuicListen(addr, port)
	n := new(typeQUIC)
	n.l = *lis
	return n
}
