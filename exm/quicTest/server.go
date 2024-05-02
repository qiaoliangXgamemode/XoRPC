package main

import (
	"context"
	"crypto/tls"
	"github.com/quic-go/quic-go"
	"fmt"
	"io"
)

func main() {
	addr := "localhost:8080"
	listener, err := quic.ListenAddr(addr, GenerateTLSConfig(), nil)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		session, err := listener.Accept(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Printf("GO ~\r\n")
		go handleSession(session)
	}
}

func handleSession(session quic.Connection) {
	stream, err := session.AcceptStream(context.Background())
	fmt.Printf("[]GO ~\r\n")
	if err != nil {
		panic(err)
	}
	// _, err = io.Copy(loggingWriter{stream}, stream)
	
	// buf 1024
	buf := make([]byte, 1024)

	// stream.SendStream.Writer.Write([]byte("test"))
	// _, err = io.Copy(loggingWriter{stream}, stream)

	_, err = io.ReadFull(stream, buf)
	
	if err != nil {
		panic(err)
	}

	fmt.Printf("Received message: %s\n", buf)
	// size 40
	io.Writer.Write(stream,[]byte("testQUICtestQUICtestQUICtestQUICtestQUIC"))
	defer stream.Close()
}

func GenerateTLSConfig() (*tls.Config) {
	cert, err := tls.LoadX509KeyPair("xianghucloud.cn_bundle.crt", "xianghucloud.cn.key")
	if err != nil {
		panic(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	config.NextProtos = []string{"quic-echo-example"}

	return config
}

// type loggingWriter struct{ io.Writer }
// func (w loggingWriter) Write(b []byte) (int, error) {
// 	fmt.Printf("Server: Got '%s'\n", string(b))
// 	return w.Writer.Write(b)
// }