package XoRPC

import (
	"XoRPC/XoRPC/cfg"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/xtaci/kcp-go"
)

// Listen
func (node *NodeXoRPC) Nodehandle(nodeListen interface{}) {
	switch v := nodeListen.(type) {
	case *net.Conn:
		node.nodehandleUDP(v)
	case *kcp.Listener:
		node.nodehandleKCP(v)
	default:
		fmt.Println("Unsupported type returned")
	}
}

func (node *NodeXoRPC) PnodehandleUDP(listener *net.UDPConn) {
	for {
		data := make([]byte, 1024)
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			log.Printf("[Info][handle][PnodehandleUDP] %s", err)
		}
		log.Printf("[Info][handle][Node][Conn] addr=%s", remoteAddr.String())
		if Mapfomain, ok := cfg.ParesJsonCONFIG(string(data[n:])); ok == true {
			if Mapfomain.Type == "Node" && Mapfomain.Version == Node_Version {
				res := cfg.POutNodeRespon(
					time.Now().String(),
					sureNode_Version,
					node.config.ServiceName,
					node.config.ServiceID,
				)
				// Send respon, after keeplive
				listener.WriteTo([]byte(res), remoteAddr)
				// Find NodeIndex APPHash
				if val, ok := node.PNodeFindIndex(Mapfomain.APPHash); ok {
					go node.PNodePushIndex(val)
				} else {
					node.PAddNodeIndex(node.PSetNodeIndex(Mapfomain.AppNodeID,
						remoteAddr.String(),
						"UDP",
						Mapfomain.AppName,
						Mapfomain.APPHash,
						remoteAddr))
				}
			}
		}
	}
}

func (node *NodeXoRPC) PnodehandleKCP(listener *kcp.Listener) {

}

func (node *NodeXoRPC) PnodehandleQUIC() {

}

// Conn
func (node *NodeXoRPC) PNodehandle(nodeListen interface{}) {
	switch v := nodeListen.(type) {
	case *net.Conn:
		node.nodehandleUDP(v)
	case *kcp.Listener:
		node.nodehandleKCP(v)
	default:
		fmt.Println("Unsupported type returned")
	}
}

func (node *NodeXoRPC) nodehandleUDP(Conn *net.Conn) {

}

func (node *NodeXoRPC) nodehandleKCP(Conn *kcp.Listener) {

}
