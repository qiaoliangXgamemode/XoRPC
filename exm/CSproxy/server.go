package main

import(
	"fmt"
	"io"
    "net"
	"github.com/xtaci/kcp-go"
)

func main() {
	kcps, err := kcp.ListenWithOptions("127.0.0.1:10000", nil, 10, 3) // kcp
	fmt.Printf("/kcp/127.0.0.1/10000\r\n")
	// tcps, err := net.Listen("tcp", ":20000") // tcp
	if err != nil {
		panic(err)
	}
	
	conn, e := kcps.AcceptKCP()
	if e != nil {
		panic(e)
	}
	
	hlash(conn)
	// _, err = io.ReadFull(conn, body)
	// n, err := conn.Read(body)
}

func hlash(conn net.Conn) {
	body := make([]byte, 1024)
	n, e := conn.Read(body)
	if e != nil {
		panic(e)
	}
	// fmt.Printf(string(body[n-4:n]), body[n-4:n])
	for {
		gohaool(conn, string(body[n-4:n]))
	}
}

func gohaool(k net.Conn, addres string) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "127.0.0.1", addres))
	if err != nil {
		panic(err)
	}
	fmt.Printf(lis.Addr().String())
	fmt.Printf("/tcp/127.0.0.1/"+addres + "\r\n")
	conn, err := lis.Accept()
	if err != nil {
		panic(err)
	}
	go io.Copy(conn, k)
	go io.Copy(k, conn)
}