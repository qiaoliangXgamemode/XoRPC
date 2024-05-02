package main

import (
	"fmt"
	"net"
)

func main() {
	// STUN服务器地址和端口
	stunServer := "stun.qq.com:19302"

	// 创建TCP连接
	conn, err := net.Dial("tcp", stunServer)
	if err != nil {
		fmt.Println("无法连接到STUN服务器:", err)
		return
	}
	defer conn.Close()

	// STUN请求消息
	request := []byte{
		0x00, 0x01, 0x00, 0x00, // Message Type: Binding Request
		0x21, 0x12, 0xA4, 0x42, // Magic Cookie
		0x78, 0x65, 0x00, 0x00, // Transaction ID
	}

	// 发送STUN请求
	_, err = conn.Write(request)
	if err != nil {
		fmt.Println("发送STUN请求失败:", err)
		return
	}

	// 接收STUN响应
	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil {
		fmt.Println("接收STUN响应失败:", err)
		return
	}

	// 解析XOR-Mapped-Address
	xorMappedAddress := parseXORMappedAddress(response[:n])
	fmt.Println("XOR-Mapped-Address:", xorMappedAddress)
}

// 解析XOR-Mapped-Address
func parseXORMappedAddress(data []byte) string {
	// STUN消息头部长度为20字节
	// XOR-Mapped-Address属性的类型为0x0020，长度为8字节
	// XOR-Mapped-Address属性的值从第28字节开始
	ip := net.IP{
		data[28] ^ data[3],
		data[29] ^ data[2],
		data[30] ^ data[1],
		data[31] ^ data[0],
	}
	port := uint16(data[26])<<8 | uint16(data[27])
	return fmt.Sprintf("%s:%d", ip.String(), port)
}