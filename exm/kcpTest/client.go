package main

import (
	"fmt"
	"log"
	"net"

	// "encoding/binary"
	// "time"
	// "io"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/xtaci/kcp-go"
)

func main() {
	sraddr := ":45614"
	raddr := "150.158.94.134:2500"
	srcAddr, err := net.ResolveUDPAddr("udp", sraddr) // 45614
	dstAddr, err := net.ResolveUDPAddr("udp", raddr)  // 2333
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		panic(err)
	}
	// send
	if _, err = conn.Write([]byte("hello, I'm new peer")); err != nil {
		log.Panic(err)
	}
	data := make([]byte, 1024)

	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		log.Printf("error during read: %s", err)
	}
	conn.Close()
	log.Printf("local:%s server:%s another:%s\n", srcAddr, remoteAddr, string(data[:n]))
	PeerHole(srcAddr, string(data[:n]))
}

func PeerHole(srcAddr *net.UDPAddr, peer string) {
	var convid uint32
	convid = 513214561
	conn, err := net.ListenUDP("udp", srcAddr)
	dstAddr, err := net.ResolveUDPAddr("udp", peer) // 2333
	if err != nil {
		panic(err)
	}
	c, err := kcp.NewConn3(convid, dstAddr, nil, 10, 3, conn)
	if err != nil {
		panic(err)
	}
	c.SetACKNoDelay(false)
	go func(c *kcp.UDPSession) {
		for {
			//
			rand.Seed(time.Now().UnixNano())  //初始化种子
			b := make([]byte, 12)             //随机生成字符数组
			rand.Read(b)                      //整合
			rand_str := hex.EncodeToString(b) //转换为string
			//

			c.Write([]byte(rand_str))
			// data := 1024
			buf := make([]byte, 1024)
			// time.start
			startTime := time.Now()
			_, err := c.Read(buf)
			// time.over
			cost := time.Since(startTime) / time.Millisecond
			fmt.Printf("[Info][Node][Ping-Pong] Mssage Delay time %dms\r\n", cost)
			// _, err := io.ReadFull(c, buf)
			if err != nil {
				panic(err)
			}
			// log.Println("recv: ",string(buf[:n]))
		}
	}(c)
	for {

	}
}
