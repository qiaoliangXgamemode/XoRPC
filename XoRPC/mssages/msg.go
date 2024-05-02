package mssages

import (
	"net"
	"errors"
	"io"
)

type ServerMsgPool struct {
	conn *net.UDPConn
	rAddr *UDPAddr
	AES256 string
}

type NodeMsgPool struct {
	conn *net.UDPConn
	AES256 string
	token string
}

func PeerWriteFull(w io.Writer, buf []byte) (err error) {
	want := len(buf)
	n := 0
	for n < want && err == nil {
		var nn int
		nn, err = w.Write(buf[n:])
		n += nn
	}
	if n == want {
		err = nil
	} else {
		err = fmt.Errorf("write full n(%d) not match want(%d)", n, want)
	}
	return
}

// 建立起新的AES256私密加密
func NewAES256network(tk string, AEShash string, conn *net.UDPConn) {
	APPHash, err := WriteAES256(AEShash, conn)
	token, err := NodeApplytoken(tk, conn)
	if err != nil {
		log.Fatal(err)
	}
	p := new(NodeMsgPool)
	p.conn = conn
	p.AES256 = APPHash
	return p
}

func JsonPageares(conn *net.UDPConn, Json string) (Json string) {
	conn.Write([]byte(Json))
	data := make([]byte, 1472)  // len(MUT 1492) Or len(AES Length)
    n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
        log.Printf("[Info][Control][ERROR][JsonPageares] %s", err)
    }
	return data[:n]
}

