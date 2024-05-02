package main
import(
	"XoRPC/XoRPC/src/cfg"
	"XoRPC/XoRPC/service"
	// "log"
)

func main(){
	// 创建配置
	var Config cfg.AppConfig
	// 节点名称
	Config.AppName = "Node1"
	// 节点 ID
	Config.AppNodeID = 1
	// Token Setting
	Config.Token = "wc"
	// 节点Hash值
	Config.APPHash = "123456"
	// 相邻的节点
	// Config.NearNode = map[string]string{"Hash":"IP:Port","Hash1":"IP:Port1"}
	// 绑定地址
	Config.NodeIPV4 = "0.0.0.0"
	// 通信协议
	Config.Protocol = "KCP"
	// 绑定端口 (UDP)
	Config.SerNodePort = 10086
	// 启用转发
	Config.TranspondForwar = true
	// 注册配置
	Node := XoRPC.NewNodeXORPC(Config)
	// 输出当前节点全部信息
	// log.Printf(XoRPC.LogsConfigNode(Node))
	XoRPC.LogsConfigNode(Node)
	Node.NodeRun(4)
	// 限制转发人数，8人。（待会改，服务器用的接口）
	Node.SetTranspondForwardLimit(8)
	// 节点限制点对点人数
	Node.AddPeerLimit()
}


