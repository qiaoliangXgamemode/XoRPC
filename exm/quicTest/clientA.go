package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/quic-go/quic-go"
	"io"
)

func main() {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}
	addr := "localhost:8080"
	conn, err := quic.DialAddr(context.Background(), addr, tlsConf, nil)
	if err != nil {
		panic(err)
	}


	stream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		panic(err)
	}

	// size 1108
	message := "QUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TESTQUIC-TEST"
	// size 8
	// message := "QUIC-TEST" 

	fmt.Printf("Client: Sending '%s'\n", message)
	io.Writer.Write(stream,[]byte(message))
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)
	
	_, err = io.ReadFull(stream, buf)
	fmt.Printf("Received message: %s\n", buf)
	if err != nil {
		panic(err)
	}
	// defer conn.Close()
	
	// send
	// for{
			
	// 	go func(stream quic.Stream){

	// 		message := "GO-QUIC-TEST"
	// 		fmt.Printf("Client: Sending '%s'\n", message)
	// 		_, err = stream.Write([]byte(message))
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		buf := make([]byte, 1024)
	// 		_, err = io.ReadFull(stream, buf)
	// 		fmt.Printf("Received message: %s\n", buf)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}(stream)
	// }
}