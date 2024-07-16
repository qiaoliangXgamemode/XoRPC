package XoRPC

func (cfg *NodeConfig) UpdateCfg(cfgs *NodeConfig) {
}

func (cfg *NodeConfig) Run() {
	network_do(cfg)
}
