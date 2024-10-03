package nat

import (
	"errors"
)

var (
	// AES KEY
	AES256                        = "0123456789abcdef0123456789abcdef" // fixed encryption keys
	JinZhiShiYongWeiBeiAnDeSheBei = errors.New("未经备案不允许运行")
	// Error
	ErrInvalidParam = errors.New("Incorrect network address,Please check NodeIPV4/6 and SerNodePort")
	ErrNetDail      = errors.New("Incorrect network address, Dial Addres error.")
	ErrNetnil       = errors.New("Incorrect network address, Addres error.")
)

// return Returns a different protocol.
// (*Listen) udp or kcp, default udp
func SiNodeListen(mode string, addres string) (listen interface{}, e error) {
	switch mode {
	case "UDP":
		return CreateUDPListen(addres)
	case "KCP":
		return CreateKCPListen(addres)
	default:
		return CreateUDPListen(addres)
	}
	return nil, ErrNetnil
}

// return Returns a different protocol.
// (*Conn) udp or kcp, default udp
func SiNodeDial(mode string, addres string) (dail interface{}, e error) {
	switch mode {
	case "UDP":
		return CreateUDPConn(addres)
	case "KCP":
		return CreateKCPConn(addres)
	default:
		return CreateUDPConn(addres)
	}
	return nil, ErrNetnil
}

// return *Listen Protocol
// (*Listener) udp or kcp, default udp
func LisNodeDial(mode string, addres string) (dail interface{}, e error) {
	switch mode {
	case "UDP":
		return CreateUDPListen(addres)
	case "KCP":
		return CreateKCPListen(addres)
	default:
		return CreateUDPListen(addres)
	}
	return nil, ErrNetnil
}
