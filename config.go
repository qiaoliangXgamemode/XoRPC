package XoRPC

const (
	// version number
	Node_Version = "0.0.1"
	// 子节点
	OneNode_Version = "0.0.1"
)

// Distributed Server cluster Node
type NodeConfig struct {
	// Node ID, This is a virtual id
	// If there is LP plugin management, this is completely managed by LP
	// LP plugin is CN service special version
	ServiceID int
	// the Node Name, This is a virtual name
	ServiceName string
	// Node weight , This is a virtual
	Serviceweight int
	// 通信加密 由AES256和TLS，一般来说QUIC协议默认启用加密

	// encryption enable
	// encryption AES256 and TLS
	ServiceEncrypt bool

	// 过滤器，可为HTTP、HTTPS、Minecraft(IntVar)达到过滤
	// 写入XoRPC是为LP插件使用

	// Filter enable
	ServiceFilter bool
	// Filter type... select HTTP、HTTPS、Minecraft(IntVar)?
	ServiceFiltertype string

	// 公开密钥 与 私密密钥
	// 公开密钥作用是在公域网上被其他节点验证，并且进行通信
	// 私密密钥作用是在广域网上被其他节点验证，并且进行通信与监控
	// 私密密钥具有监控、日志追踪、从新选举主节点作用

	// Public token vry spDimain network
	PublicToken string
	// Private token vry WideArea network
	PrivateToken string

	// Hash分布式，不过不依赖青联LP插件服务的话，完全就是没用的东西。

	// Node setting self Hash
	ServiceGroupHash int

	// 公域网 Public area network
	// 广域网 wide area network
	// spDimain 公域网作用是公开节点地址，用于节点与节点之间在公域上连接。 一般来说RPC会主动拉取主节点或者相邻节点的信息
	// WideArea广域网用于节点与节点私密进行监控、信息交换、类似与局域网。

	// The auto is controlled by DTH and can be partially managed manually
	// WideArea networks are used for node to node privacy for monitoring, information exchange, and similar to local area networks.
	// WideArea network
	Node_widearea_Control widearea
	// The auto is controlled by DTH and can be partially managed manually
	// The spDimain(public domain network) is used to expose node addresses and connect nodes on the public domain. Generally speaking, the RPC will actively pull information about the primary node or adjacent nodes
	// spDimain network
	Node_public_spDimain *Spdomain

	// protocol and IP Addres
	Protocol string
	// Your ipv4 local addres
	NodeIPV4 string
	// Your ipv6 local addres
	NodeIPV6 string
	// Node listen port, If the DTH routing table is used, ports after NAT are displayed
	SerNodePort int

	// - Node Extension -
	// Node traffic forwarding
	TranspondForwar bool
	// Node forward port
	TranspondForwarPort int
	// Node Flow Doamin
	FlowDomain map[string]int

	// Qinglian PLUS version
	// your local rules should be supported.
	// not you computer rules, it's you local. you konw?
	DomainTUNandAPT bool

	NodeNetworkV4 interface{}
	NodeNetworkV6 interface{}
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
	Node_public_spDimain map[string]int

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
