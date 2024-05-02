package main
import (
	"XoRPC/XoRPC/src/cfg"
	"XoRPC/XoRPC/service"
	// "log"
)

func main(){
	var Config cfg.ServerConfig
	// 节点ID
	Config.ServiceID = 1
	// 节点名称
	Config.ServiceName = "服务器节点1"
	// 节点权重
	Config.Serviceweight = 1
	// 节点通信加密
	Config.ServiceEncrypt = true
	// 节点过滤器
	Config.ServiceFilter = true
	// 节点过滤器类型
	// Minecraft 节点只专注于minecraft的数据
	Config.ServiceFiltertype = "minecraft"
	// 节点公开的验证密钥
	Config.PublicToken = "1bb8snhasd(Habv)"
	// 节点私密的验证密钥
	Config.PrivateToken = "G9sdb&Ubvsad0GH*Jwds2rt4t59ndc0cn+112s.Nsm234"
	// 节点Hash唯一
	Config.ServiceGroupHash = "1dsf"
	// 广域网(可选)
	Config.Node_widearea_spDimain = {
		0:"193.22.45.111"
	}
	// 公域网(可选)
	Config.Node_public_spDimain = {}
	// 绑定地址
	// Config.NodeIPV6 = "fe80::973e:1e65:a21e:c3f3"
	Config.NodeIPV4 = "0.0.0.0"
	// 通信协议
	Config.Protocol = "UDP"
	// 绑定端口 (UDP)
	Config.SerNodePort = 2333
	// Config.conn
	// 是否启用节点流量转发
	Config.TranspondForwar = 1
	// 流量转发端口绑定 （TCP） | 可以分开UDP和TCP
	Config.TranspondForwarPort = 10086
	XoServerNode := XoRPC.NewServerXORPC(Config)
	XoRPC.LogsConfigServer(XoServerNode)
	XoRPC.NodeRun("node")
}