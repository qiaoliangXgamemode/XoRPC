package XoRPC

import (
	"XoRPC/XoRPC/nat"
	"XoRPC/XoRPC/src/cfg"
	"fmt"
	"log"
)

const (
	AES256 = "0123456789abcdef0123456789abcdef" // fixed encryption keys
)

//
// func (node *NodeXoRPC) PeerUeued(conn *net.UDPAddr) {

// }

func LogsCgfJsonNode(node *NodeXoRPC) string {
	return cfg.OutputCgfnodeJson(
		node.config.AppName,
		node.config.APPHash,
		node.config.Token,
		node.config.AppNodeID,
		node.config.NodeIPV4,
		node.config.NodeIPV6,
		node.config.TranspondForwar,
	)
}

func (node *NodeXoRPC) NodeListen() {
	srcAddr := fmt.Sprintf("%s:%s", node.config.NodeIPV4, node.config.SerNodePort)
	// Jian Ting Duan Kou
	if nodeListen, err := nat.SiNodeListen(node.config.Protocol, srcAddr); err == nil {
		node.Nodehandle(nodeListen)
	} else {
		log.Fatal("[Info][Control][NodeListen][ERROR] Listen Protocol error: %s", err)
	}

}

func (node *NodeXoRPC) PnodeListen() {
	srcAddr := fmt.Sprintf("%s:%s", node.config.NodeIPV4, node.config.SerNodePort)
	nat.SiNodeDial(node.config.Protocol, srcAddr)
}
