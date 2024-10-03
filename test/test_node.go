package main

import (
	"XoRPC"
	"math/rand"
	"time"
)

func main() {
	var cfgs XoRPC.NodeConfig
	cfgs.ServiceID = 1001
	cfgs.ServiceName = "Minecraft Connent Node"
	cfgs.Serviceweight = 5
	cfgs.ServiceEncrypt = true
	cfgs.PublicToken = "让你进来吗臭傻逼"
	cfgs.PrivateToken = "1145145"
	cfgs.Protocol = "QUIC"
	cfgs.NodeIPV4 = "0.0.0.0"
	cfgs.SerNodePort = 2500
	cfgs.ServiceFilter = true
	cfgs.ServiceFiltertype = "Minecraft"
	// cfgs.Initialize()
	// cfgs.Node_public_spDimain
	// cfgs.Node_public_spDimain.AddrouteSpNode(1, "127.0.0.1", 2500, "Node1", 2)
	cfgs.Run()
	cfgs.Node_public_spDimain.AddrouteSpNode(1, "240e:350:75a5:f00:dc9a:a28:c8c8:a34d", 2500, "Node8", 2)
	for {

	}

}

func RandomSizeTen() string {
	rand.Seed(time.Now().UnixNano())
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
