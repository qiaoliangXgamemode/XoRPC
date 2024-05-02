package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	localAddr := "127.0.0.1:8080"     // 本地监听地址
	remoteAddr := "127.0.0.1:25565"   // 远程服务器地址

	listener, err := net.Listen("tcp", localAddr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Listening on %s\n", localAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr())

		go handleConnection(conn, remoteAddr)
	}
}

func handleConnection(localConn net.Conn, remoteAddr string) {
	remoteConn, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		fmt.Println("Error connecting to remote:", err)
		localConn.Close()
		return
	}
	// defer remoteConn.Close()

	fmt.Printf("Connected to remote %s\n", remoteAddr)

	// 启动两个 goroutine 分别进行数据转发
	go copyData(localConn, remoteConn)
	go copyData(remoteConn, localConn)
}

func copyData(dst io.WriteCloser, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("Error copying data:", err)
	}
	defer dst.Close()
}
