package XoRPC

import (
	"context"
	"log/slog"

	"github.com/quic-go/quic-go"
)

// IPV6
func (cfg *NodeConfig) quic_accept(listen *typeQUIC) {
	for {
		ss, e := listen.l.Accept(context.Background())
		if e != nil {
			slog.Info("[IPV6][Stack][quic_accept] QUIC error", e)
		}
		// add ipv6 proxy local
		cfg.add_xm_porxy(1, ss, "")
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
