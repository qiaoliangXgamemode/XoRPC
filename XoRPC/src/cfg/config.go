package cfg

import(

)

// Distributed Server cluster Node
type ServerConfig struct {
	EnabledNode int // 是否启用中转节点
	ControlID int // 监控ID
	ServerName string
	ServerNodeID int
	ServerGroupHash string // 监控密钥
	 // 广域网 集群服务器白名单 | 质量:低
	 // spDimain
	Server_Control_Whitelist map[string]int

	// IP info
	NodeIPV4 string
	NodeIPV6 string // Add later
	
	Protocol string
	SerNodePort int


	// --- Server Extension ---
	// TCP Forwarding stream
	TranspondForwar int
	TranspondForwarPort int
	// 中转网 | 质量优先: 高
	// flowDomain
	// FlowDomain map[string]int
}

// Distributed node
type AppConfig struct {
	AppName string
	AppNodeID int
	APPHash string
	// 2^8 AddrsHash
	NearNode map[string]string
	// pppDomain
	AppPPPdomain map[string]int // 点对网 | 质量优先: 中

	// IP info
	NodeIPV4 string
	NodeIPV6 string // Add later

	Protocol string
	SerNodePort int


	// --- App Extension ---
	// APP TCP Forwarding stream
	TranspondForwar int
	ApploctForwarAddress []string

}