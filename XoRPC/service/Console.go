package XoRPC

import (
	"log"
	"XoRPC/XoRPC/src/cfg"
)

// Node
func LogsConfigNode(node *NodeXoRPC) {
	log.Printf("[Info][Node][ID][%d] APPHash: %s", node.config.AppNodeID, node.config.Token)
	log.Printf("[Info][Node][ID][%d] AppName: %s", node.config.AppNodeID, node.config.AppName)
	log.Printf("[Info][Node][ID][%d] AppNodeID: %d", node.config.AppNodeID, node.config.AppNodeID)
	log.Printf("[Info][Node][ID][%d] APPHash: %s", node.config.AppNodeID, node.config.APPHash)
	log.Printf("[Info][Node][ID][%d] NearNode: %+v", node.config.AppNodeID, node.config.NearNode)
	log.Printf("[Info][Node][ID][%d] AppPPPdomain: %+v", node.config.AppNodeID, node.config.AppPPPdomain)
	log.Printf("[Info][Node][ID][%d] NodeIPV4: %s", node.config.AppNodeID, node.config.NodeIPV4)
	log.Printf("[Info][Node][ID][%d] NodeIPV6: %s", node.config.AppNodeID, node.config.NodeIPV6)
	log.Printf("[Info][Node][ID][%d] Protocol: %s", node.config.AppNodeID, node.config.Protocol)
	log.Printf("[Info][Node][ID][%d] SerNodePort: %d", node.config.AppNodeID, node.config.SerNodePort)
	log.Printf("[Info][Node][ID][%d] TranspondForwar: %d", node.config.AppNodeID, node.config.TranspondForwar)
	log.Printf("[Info][Node][ID][%d] ApploctForwarAddress: %+v", node.config.AppNodeID, node.config.ApploctForwarAddress)
}

// Server
func LogsConfigServer(node *ServerXoRPC) {
	log.Printf("[Info][Control][ID][%d] EnabledNode: %d", node.config.ControlID, node.config.EnabledNode)
	log.Printf("[Info][Control][ID][%d] ControlID: %d", node.config.ControlID, node.config.ControlID)
	log.Printf("[Info][Control][ID][%d] ServerName: %s", node.config.ControlID, node.config.ServerName)
	log.Printf("[Info][Control][ID][%d] ServerNodeID: %d", node.config.ControlID, node.config.ServerNodeID)
	log.Printf("[Info][Control][ID][%d] ServerGroupHash: %+v", node.config.ControlID, node.config.ServerGroupHash)
	log.Printf("[Info][Control][ID][%d] Server_Control_Whitelist: %+v", node.config.ControlID, node.config.Server_Control_Whitelist)
	log.Printf("[Info][Control][ID][%d] NodeIPV4: %s", node.config.ControlID, node.config.NodeIPV4)
	log.Printf("[Info][Control][ID][%d] NodeIPV6: %s", node.config.ControlID, node.config.NodeIPV6)
	log.Printf("[Info][Control][ID][%d] Protocol: %s", node.config.ControlID, node.config.Protocol)
	log.Printf("[Info][Control][ID][%d] SerNodePort: %d", node.config.ControlID, node.config.SerNodePort)
	log.Printf("[Info][Control][ID][%d] TranspondForwar: %d", node.config.ControlID, node.config.TranspondForwar)
	log.Printf("[Info][Control][ID][%d] TranspondForwarPort: %+v", node.config.ControlID, node.config.TranspondForwarPort)
}
