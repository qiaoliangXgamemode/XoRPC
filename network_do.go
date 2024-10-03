package XoRPC

import (
	"log/slog"
	"net"
	"strings"
)

func (cfg *NodeConfig) Network_do() {
	switch IsIPV64(cfg.NodeIPV4) {
	// case 9: CN 自主研发
	case 6:
		cfg.networkNetNode(cfg.NodeIPV6)
	case 4:
		cfg.networkNetNode(cfg.NodeIPV4)
	default:
		slog.Error("ERROR message", "Listen QUIC error", IPADDRES_EEROR)
	}
}

func (cfg *NodeConfig) networkNetNode(NodeIP string) {
	switch cfg.Protocol {
	case "QUIC":
		l := ListenerQUIC(NodeIP, cfg.SerNodePort)
		cfg.acceptStackQUIC(l)
	case "TCP":
		l := ListenerTCP(NodeIP, cfg.SerNodePort)
		cfg.acceptStackTCP(l)
	default:
		l := ListenerQUIC(NodeIP, cfg.SerNodePort)
		cfg.acceptStackQUIC(l)
	}
}

// IF isIP v4 v6
func IsIPV64(NodeIP string) int {
	if v4 := isIPv4(NodeIP); v4 == true {
		return 4
	}

	if v6 := isIPv6(NodeIP); v6 == true {
		return 6
	}
	return 0
}
func isIPv4(ipv4Addres string) bool {
	ip := net.ParseIP(ipv4Addres)
	return ip != nil && strings.Contains(ipv4Addres, ".")
}

func isIPv6(ipv6Addres string) bool {
	ip := net.ParseIP(ipv6Addres)
	return ip != nil && strings.Contains(ipv6Addres, ":")
}
