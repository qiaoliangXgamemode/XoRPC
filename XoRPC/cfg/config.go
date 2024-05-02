package cfg

import "net"

// Distributed Server cluster Node
type ServerConfig struct {
	ServiceID      int // 服务ID
	ServiceName    string
	Serviceweight  int
	ServiceEncrypt bool // encryption 通信加密

	ServiceFilter     bool   // Filter on / off
	ServiceFiltertype string // Filter type...

	// 公开密钥 与 私密密钥
	// 公开密钥作用是在公域网上被其他节点验证，并且进行通信
	// 私密密钥作用是在广域网上被其他节点验证，并且进行通信与监控
	// 私密密钥具有监控、日志追踪、从新选举主节点作用
	PublicToken  string // Public token vry
	PrivateToken string // Private token vry spDimain

	ServiceGroupHash string // Node setting self Hash
	// 公域网、广域网
	// spDimain公域网作用是公开节点地址，用于节点与节点之间在公域上连接。 一般来说RPC会主动拉取主节点或者相邻节点的信息
	// WideArea广域网用于节点与节点私密进行监控、信息交换、类似与局域网。
	// WideArea
	Node_widearea_Control map[string]int
	// spDimain
	Node_public_spDimain interface{}

	// protocol and IP Addres
	Protocol    string
	NodeIPV4    string
	NodeIPV6    string
	SerNodePort int

	// --- Server Extension ---
	TranspondForwar     bool
	TranspondForwarPort int
	FlowDomain          map[string]int

	// Qinglian PLUS version
	// your local rules should be supported.
	// not you computer rules, it's you local. you konw?
	DomainTUNandAPT bool
}

type NodeIndex struct {
	ServiceID   int
	ServiceName string
	APPHash     string
	Protocol    string
	Address     *net.UDPAddr
}

// Distributed node
type AppConfig struct {
	ServiceID      int // 服务ID
	ServiceName    string
	Serviceweight  int
	ServiceEncrypt bool // encryption 通信加密

	ServiceFilter     bool   // Filter on / off
	ServiceFiltertype string // Filter type...

	PublicToken  string // Public token vry
	PrivateToken string // Private token vry spDimain

	ServerGroupHash string // Node setting self Hash
	// 公域网、广域网
	// spDimain公域网作用是公开节点地址，用于节点与节点之间在公域上连接。
	// WideArea广域网用于节点与节点私密进行监控、信息交换、类似与局域网。
	// WideArea
	Node_widearea_spDimain map[string]int
	// spDimain
	Node_Control_spDimain map[string]int
	// public Domain
	Node_public_spDimain map[string]NodeIndex

	// protocol and IP Addres
	Protocol    string
	NodeIPV4    string
	NodeIPV6    string
	SerNodePort int

	// --- Server Extension ---
	TranspondForwar     bool
	TranspondForwarPort int
	FlowDomain          map[string]int

	// Qinglian PLUS version
	// your local rules should be supported.
	// not you computer rules, it's you local. you konw?
	DomainTUNandAPT bool
}
