package main

import "XoRPC"

func main() {
	cfgs := new(XoRPC.NodeConfig)
	cfgs.ServiceID = 1001
	cfgs.ServiceName = "Minecraft Connent Node"
	cfgs.Serviceweight = 5
	cfgs.ServiceEncrypt = true
	cfgs.PublicToken = "让你进来吗臭傻逼"
	cfgs.PrivateToken = "1145145"
	cfgs.Protocol = "QUIC"
	cfgs.NodeIPV4 = "::"
	cfgs.SerNodePort = 11451
	cfgs.ServiceFilter = true
	cfgs.ServiceFiltertype = "Minecraft"
	cfgs.Run()
}
