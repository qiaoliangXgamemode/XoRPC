package XoRPC

import (
	"XoRPC/XoRPC/cfg"
	"net"
	// "XoRPC/XoRPC/src/filter"
)

const (
	// version number
	// R -> FlowNode
	// N -> Node
	// Q -> Qinglian Branch versions, but This is Xorpc. not Qinglian
	FlowNode_Version = "0.0.1R"
	Node_Version     = "0.0.1N"
	sureNode_Version = "0.0.1Q" // Qianlian PLUS TUN.
)

// Server XoRPC
type ServerXoRPC struct {
	config cfg.ServerConfig
	// filter filter.FilterNumber // 限流水桶令
}

// Node PTP RPC
type NodeXoRPC struct {
	config cfg.AppConfig
}

type NodeIndex struct {
	ServiceID   int
	ServiceName string
	APPHash     string
	Protocol    string
	Addr        *net.UDPAddr
}

func NewServerXORPC(cfg cfg.ServerConfig) *ServerXoRPC {
	xo := new(ServerXoRPC)
	xo.config = cfg
	return xo
}

func NewNodeXORPC(cfg cfg.AppConfig) *NodeXoRPC {
	xo := new(NodeXoRPC)
	xo.config = cfg
	return xo
}

// Node Hash
func SetAppHash(node *NodeXoRPC) string {
	return node.config.ServiceName
}

func (node *ServerXoRPC) Getsize64Hash() string {
	return "test"
}

// WTF What are you doing?
func (node *NodeXoRPC) NodeRun(into string) {
	switch into {
	case "node":
		node.NodeListen()
	case "Pnode":
		node.PnodeListen()
	default:
		node.NodeListen()
	}
}

// // // // // // //
// 疯 狂 星 期 四 //
//  V  我  5  0  //
// // // // // // //

// ```
// K F C Crazy Thursday
//         V me 50
// ```
