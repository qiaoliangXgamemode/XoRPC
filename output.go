package XoRPC

import (
	"log/slog"
	"time"
)

func (cfg *NodeConfig) OutServiceALL() {
	// logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetLogLoggerLevel(slog.LevelDebug)
	cfg.OutServiceRun()
	cfg.OutServiceID()
	cfg.OutServiceName()
	cfg.OutServiceweight()
	cfg.OutServiceEncrypt()
	cfg.OutServiceFilter()
	cfg.OutServiceProtocol()
	cfg.OutServiceNodeIPV4()
	cfg.OutServiceNodeIPV6()
	cfg.OutServiceSerNodePort()
	cfg.OutServicePublicToken()
}

func (cfg *NodeConfig) OutServiceRun() {
	slog.Info("info message [Control][OutPut]", "Node Run", time.Now().String())
}

func (cfg *NodeConfig) OutServiceID() {
	slog.Info("info message [Control][OutPut]", "Node ID", cfg.ServiceID)
}

func (cfg *NodeConfig) OutServiceName() {
	slog.Info("info message [Control][OutPut]", "Node Name", cfg.ServiceName)
}

func (cfg *NodeConfig) OutServiceweight() {
	slog.Info("info message [Control][OutPut]", "Node weight", cfg.Serviceweight)
}

func (cfg *NodeConfig) OutServiceEncrypt() {
	slog.Info("info message [Control][OutPut]", "Node Name", cfg.ServiceEncrypt)
}

func (cfg *NodeConfig) OutServiceFilter() {
	slog.Info("info message [Control][OutPut]", "Node Filter", cfg.ServiceFilter)
}

func (cfg *NodeConfig) OutServiceProtocol() {
	slog.Info("info message [Control][OutPut]", "Node Protocol", cfg.Protocol)
}

func (cfg *NodeConfig) OutServiceNodeIPV4() {
	slog.Info("info message [Control][OutPut]", "Node NodeIPV4", cfg.NodeIPV4)
}

func (cfg *NodeConfig) OutServiceNodeIPV6() {
	slog.Info("info message [Control][OutPut]", "Node NodeIPV6", cfg.NodeIPV6)
}

func (cfg *NodeConfig) OutServiceSerNodePort() {
	slog.Info("info message [Control][OutPut]", "Node SerNodePort", cfg.SerNodePort)
}

func (cfg *NodeConfig) OutServicePublicToken() {
	slog.Debug("info message [Control][OutPut]", "Node OutServicePublicToken", cfg.OutServicePublicToken)
}
