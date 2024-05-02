package main

import (
	"fmt"
	"XoRPC/exm/sssssss"
)

func main() {
	connection1 := sssssss.CreateTCPListen("127.0.0.1:45555")
	for {
		conn, err := connection1.Accept()
		if err != nil {
			fmt.Println("conn ",err)
		}
		body := make([]byte, 1024)
		text := `HTTP/1.1 200 OK\r\n\r\n <h1>Chen Zihao!</h1>`
		copy([]byte(text), body)
		conn.Write(body)
		conn.Close()
	}
}