package nat

import (
	"log"
	"net"

	"github.com/xtaci/kcp-go"
)

// Peer KCP
func KCPNewPeer(conn *net.UDPConn, RemoteAddres string) *kcp.UDPSession {

	var convid uint32
	convid = 513216541 // 懒癌了
	dstAddr, err := net.ResolveUDPAddr("udp", conn.LocalAddr().String())
	conn.Close()
	c, err := net.ListenUDP("udp", dstAddr)
	if err != nil {
		log.Printf("[Info][Control][ERROR][] %s", err)
	}
	log.Printf("[Info][Control][ERROR][] %s", c.LocalAddr())
	kcpC, err := kcp.NewConn3(convid, dstAddr, nil, 10, 3, conn)
	if err != nil {
		panic(err)
	}
	return kcpC
}

func KCPtoTcpAsUDP(Kcps *kcp.UDPSession) {

}
