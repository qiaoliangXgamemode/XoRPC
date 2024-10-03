package XoRPC

var stack = false

func (cfg *NodeConfig) UpdateCfg(cfgs *NodeConfig) {
	cfg = cfgs
	cfg.Initialize()
	cfg.Node_public_spDimain.remoteAddr, _ = getRemoteIP()
}

func (cfg *NodeConfig) Run() {
	cfg.Initialize()
	go cfg.Network_do()
}

func (cfg *NodeConfig) Initialize() {
	cfg.OutServiceALL()
	cfg.Node_public_spDimain = NewSpdomain()
}
