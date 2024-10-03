package XoRPC

import (
	"context"
	"log/slog"

	"github.com/gorilla/websocket"
	"github.com/quic-go/quic-go"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// IPV6
func (cfg *NodeConfig) acceptStackQUIC(listen *protocolQUIC) {
	slog.Info("info message", "Node Accept stuat", "OK")
	// go func(l quic.Listener) {
	// 	for {
	// 		ss, e := listen.l.Accept(context.Background())
	// 		sp, _ := ss.AcceptStream(context.Background())
	// 		if e != nil {
	// 			slog.Info("[IPV6][Stack][quic_accept] QUIC error", e)
	// 		}

	// 		ReadStreamQuic(sp)
	// 	}
	// }(listen.l)
	for {
		ss, e := listen.l.Accept(context.Background())
		sp, _ := ss.AcceptStream(context.Background())
		cfg.Node_public_spDimain.RecvSpNodeReturnRoute(sp)
		// mss, ok := ReadRPCcodeStreamQuic(sp)
		// slog.Info("INFO Massge", "[IPV6][Stack][quic_accept] QUIC ", mss, ok)
		if e != nil {
			slog.Info("[IPV6][Stack][quic_accept] QUIC error", e)
		}

		// add ipv6 proxy local
		// cfg.add_xm_porxy(1, ss, "")
	}
}

// TCP
func (cfg *NodeConfig) acceptStackTCP(listen *protocolTCP) {
	slog.Info("info message", "Node Accept stuat", "OK")
	// go func(l quic.Listener) {
	// 	for {
	// 		ss, e := listen.l.Accept(context.Background())
	// 		sp, _ := ss.AcceptStream(context.Background())
	// 		if e != nil {
	// 			slog.Info("[IPV6][Stack][quic_accept] QUIC error", e)
	// 		}

	// 		ReadStreamQuic(sp)
	// 	}
	// }(listen.l)
	for {
		// conn, e := listen.l.Accept()
		// cfg.Node_public_spDimain.RecvSpNodeReturnRoute(conn)
		// mss, ok := ReadRPCcodeStreamQuic(sp)
		// slog.Info("INFO Massge", "[IPV6][Stack][quic_accept] QUIC ", mss, ok)
		// if e != nil {
		// 	slog.Info("[IPV6][Stack][quic_accept] QUIC error", e)
		// }

		// add ipv6 proxy local
		// cfg.add_xm_porxy(1, ss, "")
	}
}

// IPV6
func (cfg *NodeConfig) add_xm_porxy(ncxid int, conn quic.Connection, local interface{}) *porxytunnel {
	tunnel := new(porxytunnel)
	tunnel.ncx_id = ncxid
	tunnel.conn = conn
	tunnel.porxy_conn = local
	return tunnel
}

// IPV6
func (cfg *NodeConfig) porxy_local(string) {
	// cfg.NodeNetworkV6
}
