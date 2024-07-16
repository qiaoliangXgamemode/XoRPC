package XoRPC

import (
	"net"
	"strings"
)

func network_do(cfg *NodeConfig) {
	if v4 := IsIPv4(cfg.NodeIPV4); v4 == true {
		cfg.NodeNetworkV4, _ = CreateQuicListen(cfg.NodeIPV4, cfg.SerNodePort)
	}
	if v6 := IsIPv4(cfg.NodeIPV6); v6 == true {
		cfg.NodeNetworkV6, _ = CreateQuicListen(cfg.NodeIPV6, cfg.SerNodePort)
	}
}

func IsIPv4(ipv4Addres string) bool {
	ip := net.ParseIP(ipv4Addres)
	return ip != nil && strings.Contains(ipv4Addres, ".")
}

func IsIPv6(ipv6Addres string) bool {
	ip := net.ParseIP(ipv6Addres)
	return ip != nil && strings.Contains(ipv6Addres, ":")
}
