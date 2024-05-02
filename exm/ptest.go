package main

import (
	"net"
	"log"
	"XoRPC/XoRPC/src/encrypt"
)

func main() {
	raddr := "123.207.35.211:2333"
	conn, err := net.Dial("udp", raddr)
	json := `{
		"type":"Node",
		"version":"0.0.1N",
		"Token":"Kolashiw",
		"AppName":"Windows IX",
		"APPHash":"MhZquBuoM2ApcMac",
		"ID":1,
		"Node_IPV4":"11.45.145.0",
		"Node_IPV6":"",
		"TranspondForwar": true
	}`
	if err != nil {
		panic(err)
	}
	// log.Printf("<%s> : <%s>","test",string(json))
	// AES256 := "0123456789abcdef0123456789abcdef"  // fixed encryption keys
	// encrypted, _ := encrypt.AES256Encrypt([]byte(json), []byte(AES256))
	// decrypteds, err := encrypt.AES256Decrypt([]byte(encrypted), []byte(AES256))
	// log.Printf("<%s> : <%s>","test",string(decrypteds))
	test(conn, json)
	for i := 0; i <= 50000; i++ {
		log.Printf("cout: %d",i)
		go test(conn, json)
	}
}

func test(conn net.Conn, json string) {
	
	AES256 := "0123456789abcdef0123456789abcdef"  // fixed encryption keys
	encrypted, _ := encrypt.AES256Encrypt([]byte(json), []byte(AES256))
	// decrypteds, err := encrypt.AES256Decrypt([]byte(encrypted), []byte(AES256))
	// log.Printf("<%s> : <%s>","test",string(decrypteds))
	// send
	conn.Write([]byte(encrypted)) 

	data := make([]byte, 1024)
	
	n, err := conn.Read(data)
	if err != nil {
		log.Printf("error during read: %s", err)
	}
	decrypted, _ := encrypt.AES256Decrypt([]byte(string(data[:n])), []byte(AES256))
	log.Printf("<%s> : %s","test",string(decrypted))

}