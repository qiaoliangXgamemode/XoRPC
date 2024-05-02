package XoRPC

import (
	"XoRPC/XoRPC/cfg"
	"net"
)

func (node *NodeXoRPC) PNodeFindIndex(apphash string) (Index cfg.NodeIndex, ok bool) {
	if val, ok := node.config.Node_public_spDimain[apphash]; ok {
		return val, true
	}
	return cfg.NodeIndex{}, false
}

func (node *NodeXoRPC) PSetNodeIndex(serviceID int,
	protocol string,
	serviceName string,
	apphash string,
	addr string,
	Conn *net.UDPAddr) cfg.NodeIndex {
	return cfg.NodeIndex{
		ServiceID:   serviceID,
		ServiceName: serviceName,
		APPHash:     apphash,
		Protocol:    protocol,
		Address:     Conn,
	}
}

func (node *NodeXoRPC) PAddNodeIndex(Index cfg.NodeIndex) {
	node.config.Node_public_spDimain[Index.APPHash] = Index
	go node.PNodePushIndex(node.config.Node_public_spDimain[Index.APPHash])
}

func (node *NodeXoRPC) PNodePushIndex(Index cfg.NodeIndex) {
	for {

	}
}
