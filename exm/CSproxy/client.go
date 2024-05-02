package main

import (
	"net"
	"io"
	"fmt"
	"time"
	// "net"
	"github.com/xtaci/kcp-go"
)

const(
	typeAA = 1
)

func main() {
	Addrs := "127.0.0.1:10000"
	kcpn, err := kcp.DialWithOptions(Addrs, nil, 10, 3)
	if err != nil {
		panic(err)
	}
	AddIP := "1230"
	// WriteMsg(kcpn, AddIP)
	kcpn.Write([]byte(AddIP))
	// s := ReadMsg(kcpn)
	time.Sleep(3 * time.Second)
	for {
		// PSCopys(kcpn)
	}
}

func PSCopys(dconn net.Conn) {
	conn, err := net.Dial("tcp","127.0.0.1:25565")
	if err != nil {
		panic(err)
	}
	go io.Copy(dconn, conn)
	go io.Copy(conn, dconn)
}

func ReadMsg(r net.Conn) (bytes []byte) {
	var body []byte = make([]byte, 1024)
	_, err := io.ReadFull(r, body)
	if err != nil {
		panic(err)
	}
	return body
}

func WriteMsg(w net.Conn, ProxyAddres string) {
	var l int
	var messages []byte

	msgAddr := []byte(ProxyAddres)
	l = len(messages)
	b := make([]byte, l+1)
	b[0] = typeAA
	copy(b[0:], msgAddr)
	writeFull(w, b) // write
}

func writeFull(w io.Writer, buf []byte) (err error) {
	n, err := w.Write(buf)
	if err != nil {
		panic(err)
	}
	if n > 0 {
		fmt.Println(n)
	}
	return
}