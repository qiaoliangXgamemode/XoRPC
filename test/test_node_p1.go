package main

import (
	"XoRPC"
	"log/slog"
	"time"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	var cfgs XoRPC.NodeConfig
	cfgs.ServiceID = 1001
	cfgs.ServiceName = "Minecraft Connent Node"
	cfgs.Serviceweight = 5
	cfgs.ServiceEncrypt = true
	cfgs.PublicToken = "让你进来吗臭傻逼"
	cfgs.PrivateToken = "1145145"
	cfgs.Protocol = "QUIC"
	cfgs.NodeIPV4 = "127.0.0.1"
	cfgs.SerNodePort = 11453
	cfgs.ServiceFilter = true
	cfgs.ServiceFiltertype = "Minecraft"
	cfgs.Initialize()
	// cfgs.Node_public_spDimain

	time.Sleep(3000)
	cfgs.Node_public_spDimain.AddrouteSpNode(1, "127.0.0.1", 11451, "Node1", 2)
	cfgs.Node_public_spDimain.AddrouteSpNode(2, "8.134.186.138", 2500, "Node2", 2)
	// cfgs.Node_public_spDimain.AddrouteSpNode(2, "240e:350:75a5:f00:dc9a:a28:c8c8:a34d", 11451, "Node2", 2)
	// cfgs.Node_public_spDimain.AddrouteSpNode(3, "8.134.186.138", 2500, "Node3", 2)
	// cfgs.Node_public_spDimain.AddrouteSpNode(4, "8.134.186.138", 2500, "Node4", 2)
	// cfgs.Node_public_spDimain.FindlocalNode("Node1")
	go cfgs.Run()
	// time.Sleep(2000)
	ok := cfgs.FindNodeALL(cfgs.Node_public_spDimain, "Node8")
	if ok {
		cfgs.Node_public_spDimain.FindonlylocalNode("Node8")
	}
	for {

	}
}

// func RandomSizeTen() string {
// 	rand.Seed(time.Now().UnixNano())
// 	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	b := make([]byte, 10)
// 	for i := range b {
// 		b[i] = letterBytes[rand.Intn(len(letterBytes))]
// 	}
// 	return string(b)
// }
