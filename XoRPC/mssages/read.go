package mssages

import ( 
	"net"
)

func FromReadUDP(conn *net.UDPConn) {
	return
}

func FromReadAES256(conn *net.UDPConn, Text byte[], AES256 string) (Value string) {
	decrypted, err := encrypt.AES256Decrypt([]byte(Text), []byte(AES256))
	if err != nil {
		log.Fatal(err)
	}
	return decrypted
}

func newServerMsgPool(conn *net.UDPConn, rAddr *UDPAddr) (*ServerMsgPool) [
	p := new(ServerMsgPool)
	AppHash := remoteMoTOAES256(conn)
	p.conn = conn
	p.rAddr = remoteAddr
	p.AES256 = APPHash
	return p
]

func ReadFromUdp(conn *net.UDPConn) (*ServerMsgPool, buf string, remoteAddr *UDPAddr, err error) {
	p := newServerMsgPool(conn, remoteAddr)
	data := make([]byte, 1492)
	n, remoteAddr, err := p.conn.ReadFromUDP(data)
	if err != nil {
		log.Fatal(err)
	}
	
	return p, FromReadAES256(data[:n], node.Hash), _
}