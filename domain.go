package XoRPC

import (
	"encoding/json"
	"log/slog"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/quic-go/quic-go"
)

type FindRoute struct {
	Type     string
	Findnode string
}

type ReturnRoute struct {
	Type     string
	NULL     bool
	Findnode string
	ID       int
	Addres   string
	Port     int
	RTT      int
}

// d
type ControlDomain struct {
}

type Spdomain struct {
	nodes      map[string]nodeRoute
	remoteAddr string
	remotePort int
}

type widearea struct {
	nodes map[string]nodeRoute
}

type nodeRoute struct {
	nodeID int
	addr   string
	port   int
	weight int
	k      int
	hash   string
}

// Hash domain
// hash(IPv4)% 2^32
// but hash use ipv4 ... so we can add ServiceID(NodeID).
// only hash, Node1 Node2 ... Hash+ServiceID.
// Node sequence number conflict, ok? thank about, oh yeah. join not fair competition
// MAC Conflict? Yes
// Node1 not Node2, Node1(ServiceGroupHash) = Node2(ServiceGroupHash), every Node Don't Network!
func (cfg *NodeConfig) hashUpdate() {
	remoteAddrs := cfg.Node_public_spDimain.remoteAddr
	segments := strings.Split(remoteAddrs, ".")
	l := len(segments)
	var inhash float64
	for i := range l {
		num, _ := strconv.Atoi(segments[i])
		inhash += float64(num) * math.Pow(256, float64(l-1-i))
	}
	cfg.ServiceGroupHash = int(inhash) + cfg.ServiceID
}

func newFindRoutePong(t string, nodehash string) []byte {
	// pong := new(FindRoute)
	// pong.Type = t
	// pong.Findnode = nodehash
	pong := FindRoute{
		Type:     t,
		Findnode: nodehash,
	}
	j, _ := json.Marshal(pong)
	return j
}

func newReturnRoutePong(t string, id int, addres string, port int, nodehash string) []byte {
	// pong := new(FindRoute)
	// pong.Type = t
	// pong.Findnode = nodehash
	pong := ReturnRoute{
		Type:     t,
		NULL:     true,
		Findnode: nodehash,
		ID:       id,
		Addres:   addres,
		Port:     port,
		RTT:      rand.Intn(500), // 100ms ~ 600ms
	}
	j, _ := json.Marshal(pong)
	return j
}

// Parsing Json Context.
func passReturnRoutePong(t string) ReturnRoute {
	var p ReturnRoute
	e := json.Unmarshal([]byte(t), &p)
	if e != nil {
		slog.Debug("Debug message", "Json Unmarshal ERROR", e)
	}
	return p
}
func passFindRoutePong(t string) FindRoute {
	var p FindRoute
	e := json.Unmarshal([]byte(t), &p)
	if e != nil {
		slog.Debug("Debug message", "Json Unmarshal ERROR", e)
	}
	return p
}
func (Sp *Spdomain) RecvSpNodeReturnRoute(q quic.Stream) {
	w, code := ReadRPCcodeStreamQuic(q)
	if code == 201 {
		r := passFindRoutePong(string(w))
		if ok := Route_Type(r); ok == true {
			hashnode := FindlocalNode(Sp, r.Findnode)
			slog.Info("INFO message", "FIND ROUTE CALL hash", hashnode[r.Findnode].hash)
			id := hashnode[r.Findnode].nodeID
			addres := hashnode[r.Findnode].addr
			port := hashnode[r.Findnode].port
			hash := hashnode[r.Findnode].hash
			p := newReturnRoutePong("FINDNODERTURN", id, addres, port, hash)
			// slog.Debug("Debug message", "json reply", p)
			// slog.Debug("Debug message", "Find Node", "RetrunNodeINFO")
			WriteStreamQuic(q, p)
		}
	}

}

func (cfg *NodeConfig) FindSpNode(Hash string) {
	cfg.FindNodeALL(cfg.Node_public_spDimain, Hash)

}

func (cfg *NodeConfig) FindNodeALL(Sp *Spdomain, Hash string) bool {
	p := newFindRoutePong("FINDNODE", Hash)
	p = newFindNodeMssagePack(p)

	for _, v := range Sp.nodes {

		// slog.Debug("Debug message", "Dail Addres Route", k)
		// slog.Debug("Debug message", "Dail Addres Route IP", v.addr)
		// slog.Debug("Debug message", "Dail Addres Route Port", v.port)
		conn, e := CreateQuicConn(v.addr, v.port)
		if e != nil {
			// panic(e)
			slog.Error("[Control][CreateQuicConn][ERROR] Dail QUIC error: %s", e)
			continue
		} else {
			stream, _ := conn.OpenStream()
			rp := string(WriteStreamQuic(stream, p))
			retrunpong := passReturnRoutePong(rp)
			// slog.Debug("Debug message", "json reply", retrunpong)
			if ok := Route_Type(retrunpong); ok == true {
				cfg.Node_public_spDimain.AddrouteSpNode(retrunpong.ID, retrunpong.Addres, retrunpong.Port, retrunpong.Findnode, 2)
				return true
			}
		}
	}
	return false
}
func (Sp *Spdomain) FindonlylocalNode(Hash string) {

	slog.Debug("Debug message",
		"FindOnly [local] RouteSpNode",
		Sp.nodes[Hash].nodeID,
		"Addres",
		Sp.nodes[Hash].addr,
		"Port",
		Sp.nodes[Hash].port,
		"HashApp",
		Sp.nodes[Hash].hash)
}

func (Sp *Spdomain) FindlocalNode(Hash string) {

	slog.Debug("Debug message",
		"Findlocal RouteSpNode",
		Sp.nodes[Hash].nodeID,
		"Addres",
		Sp.nodes[Hash].addr,
		"Port",
		Sp.nodes[Hash].port,
		"HashApp",
		Sp.nodes[Hash].hash)
}

func FindlocalNode(Sp *Spdomain, Hash string) map[string]nodeRoute {
	if Sp.nodes[Hash].hash == Hash {
		return Sp.nodes
	}
	return nil
	// for k, v := Sp.nodes {
	// 	if k == Hash {

	// 	}
	// }
}

func NewSpdomain() *Spdomain {
	sp := new(Spdomain)
	sp.nodes = make(map[string]nodeRoute)
	rIP, _ := GetRemoteIP()
	sp.remoteAddr = rIP
	return sp
}

func (Sp *Spdomain) AddrouteSpNode(id int, addr string, port int, hashApp string, weight int) {
	r := new(nodeRoute)
	r.nodeID = id
	r.addr = addr
	r.port = port
	r.weight = weight
	r.k = 1
	// r.hash = RandomSizeTen()
	r.hash = hashApp
	Sp.nodes[hashApp] = *r
	slog.Debug("Debug message", "[RouteSpNode]", id, "Addres", addr, "Port", port, "HashApp", hashApp)
}

func (Sp *Spdomain) AddrouteAreaNode(id int, addr string, port int, weight int) {

}

func NodehashXOR() {

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
