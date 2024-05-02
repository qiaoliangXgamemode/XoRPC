package main

import(
	"net"
	"log"
	"time"
	// "github.com/xtaci/kcp-go"
)


func main() {
	
	
	kcps, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 2500})
	if err != nil {
		panic(err)
	
	}
	log.Println(net.UDPAddr{IP: net.IPv4zero, Port: 2500})
	log.Println(kcps.LocalAddr())
	log.Printf("[server] on")
	peers := make([]net.UDPAddr, 0, 2)
	for{
		data := make([]byte, 1024)
		n, remoteAddr, err := kcps.ReadFromUDP(data)
		log.Printf("<%s> %s\n", remoteAddr.String(), data[:n])	
		peers = append(peers, *remoteAddr)
		if err != nil {
			log.Fatal(err)
		}
		
		if len(peers) == 2 {
			log.Printf("进行UDP打洞,建立 %s <--> %s 的连接\n", peers[0].String(), peers[1].String())
			kcps.WriteToUDP([]byte(peers[1].String()), &peers[0])
			kcps.WriteToUDP([]byte(peers[0].String()), &peers[1])
			time.Sleep(time.Second * 8)
			log.Println("中转服务器退出,仍不影响peers间通信")
		}
	}
	
}

// handleEcho send back everything it received
// func handleEcho(conn *kcp.UDPSession) {
// 	for{
// 		print("go-> ",conn.RemoteAddr().String(),"\r\n")
// 		if len(peers) == 2 {
// 			// buf := make([]byte, 4096)
// 			// n, err := conn.Read(buf)
// 			// log.Println("recv: ",string(buf[:n]))
// 			// if err != nil {
// 			// 	log.Println(err)
// 			// }
// 			log.Println(conn.RemoteAddr() == peers[0].RemoteAddr())
// 			if conn.RemoteAddr() == peers[1].RemoteAddr() {
// 				conn.Write([]byte(peers[0].RemoteAddr().String()))

// 			}
// 			if conn.RemoteAddr() == peers[0].RemoteAddr() {
// 				conn.Write([]byte(peers[1].RemoteAddr().String()))

// 			}
// 		}
// 	}
// 	defer conn.Close()
// 	}