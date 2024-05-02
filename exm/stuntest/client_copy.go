package main

import (
	"fmt"
	"net"
	"log"
	"time"
	"XoRPC/XoRPC/src/cfg"
	"XoRPC/XoRPC/src/encrypt"
	"github.com/pion/stun"
)

var (
	AES256 = "0123456789abcdef0123456789abcdef"  // fixed encryption keys
)

func main() {
	sraddr := ":45614"
	raddr := "stun.l.google.com:19302"
	srcAddr,err := net.ResolveUDPAddr("udp", sraddr)  // 45614
	dstAddr, err := net.ResolveUDPAddr("udp", raddr)  // 2333
    conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}

	// 创建一个STUN消息
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)

	// 设置超时
	err = conn.SetDeadline(time.Now().Add(time.Second * 5))
	if err != nil {
		fmt.Println("Failed to set deadline:", err)
		return
	}

	// 发送STUN消息
	if _, err = conn.Write(message.Raw); err != nil {
		fmt.Println("Failed to write:", err)
		return
	}

	// 读取STUN响应
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Failed to read:", err)
		return
	}

	// 解析STUN响应
	response := new(stun.Message)
	response.Raw = buf[:n]
	if err = response.Decode(); err != nil {
		fmt.Println("Failed to decode:", err)
		return
	}

	// 获取XORMappedAddress
	var xorMappedAddress stun.XORMappedAddress
	if err = xorMappedAddress.GetFrom(response); err != nil {
		fmt.Println("Failed to get XOR-MAPPED-ADDRESS:", err)
		return
	}
	conn.Close()
	log.Printf("[Info][Control][IPV4][Addr] LocalAddrs: %s", conn.LocalAddr().String())
	fmt.Println("IP:", xorMappedAddress.IP, "Port:", xorMappedAddress.Port)
	// srcAddr,_ := net.ResolveUDPAddr("udp", conn.LocalAddr().String())
	Local, err := net.ListenUDP("udp", srcAddr )
	if err != nil {
		panic(err)
	}
	log.Printf("[Info][Control][IPV4][Addr] LocalAddr: %s", Local.LocalAddr())
	PSX := string(xorMappedAddress.IP) + ":" + string(xorMappedAddress.Port)
	srcs,_ := net.ResolveUDPAddr("udp4", PSX)
	go Local.WriteToUDP([]byte("Hello"), srcs)
	go Local.WriteToUDP([]byte("Hello"), srcs)
	go Local.WriteToUDP([]byte("Hello"), srcs)
	for{
		data := make([]byte, 1024)
		n, remoteAddr, err := Local.ReadFromUDP(data)
		if err != nil {
			log.Fatal(err)
		}
		
		decrypted, err := encrypt.AES256Decrypt([]byte(data[:n]), []byte(AES256))
		log.Printf("[Info][Control][ListenV4][Text] IP=<%s> text= %s", remoteAddr.String(), string(decrypted))
		if err != nil {
			log.Fatal(err)
		}
		// Pares Json
		if str, ok := cfg.ParesCgf(string(decrypted), remoteAddr); ok {
			encrypted, _ := encrypt.AES256Encrypt([]byte(str), []byte(AES256))
			go Local.WriteToUDP([]byte(encrypted), remoteAddr)
		}else{
			encrypted, _ := encrypt.AES256Encrypt([]byte(str), []byte(AES256))
			go Local.WriteToUDP([]byte(encrypted), remoteAddr)
		}
	}
}