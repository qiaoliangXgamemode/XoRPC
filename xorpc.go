package XoRPC

func (cfg *NodeConfig) UpdateCfg(cfgs *NodeConfig) {
	cfg = cfgs
}

func (cfg *NodeConfig) Run() {
	network_do(cfg)
}
