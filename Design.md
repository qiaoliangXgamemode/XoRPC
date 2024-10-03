# 设计
安全性，容忍性，统一标识。

# 加入网络
节点私秘钥 进行 网络寻找
公开密钥 进行 加入网络

# 加密
AES + TLS 1.2加密

# 节点算法
## 最短路径选找
dijkstra

# Logging
日志输出

# 网络分域服务 / 服务等级
flowDomain(转发域)，spDomain(广网域)，pppDomain(点对域)
flowDomain 服务质量最高
pppDomain  服务质量中等
spDomain  服务质量最低
// spDomain 不参与用户服务质量，他是一个服务器集群网络域，只负责服务器节点之间的相互通信。
# NAT
## STUN

# similar_pECnotice
监控各各节点的负载以及可用情况

# Protocol
## KCP 可靠游戏加速
KCP-GO

## QUIC
保证服务可靠性，只用于传输一些小文件与游戏加速。并不是适应与 高可靠保证 服务质量（列如：大文件传输）

## GO-RUDP
自己设计的可靠UDP(RUDP)，SR重传，保留部分ARQ。
