package cfg

import (
	"encoding/json"
	"log"
	// "XoRPC/XoRPC/service"
)

// type(
// 	Client
// )

const (
	// Code 3000
	Json_status = `{"status":"ok"}`
	// Code 2000
	Json_Error_version_not = `{"error":"The version is inconsistent with the server version!"}`
	// Code 1000
	Json_Error_unknown = `{"error":"Json Error."}`
)

// Not as good as JavaScript JSON Genshin
type PNodeRespon struct {
	Times       string `json:"times"`
	Type        string `json:"type"`
	Version     string `json:"version"`
	ServiceName string `json:"AppName"`
	ServiceID   int    `json:"ServiceID"`
}

// Not as good as JavaScript JSON Genshin
type NodeResolve struct {
	Times            string `json:"times"`
	Type             string `json:"type"`
	Version          string `json:"version"`
	PublicToken      string `json:"PublicToken"` // using PublicToken conn PNode
	AppName          string `json:"AppName"`
	APPHash          string `json:"APPHash"`
	AppNodeID        int    `json:"ID"`
	NodeIPV4         string `json:"Node_IPV4"`
	NodeIPV6         string `json:"Node_IPV6"`
	TranspondForward bool   `json:"TranspondForwar"`
}

type Peer struct {
	Types      string `json:"type"`
	NodeAddres string `json:"type"`
	RandUID    string `json:"RandUID"`
}

func ParesJsonCONFIG(Json string) (NodeResolve, bool) {
	var node NodeResolve
	err := json.Unmarshal([]byte(Json), &node)
	if err != nil {
		log.Printf("[Info][Control][ERROR][ParesCgf] ", err)
	}
	return node, true
}

// PNode respon type 200ok
func POutNodeRespon(
	times string,
	PNode_Version string,
	servicename string,
	serviceid int,
) string {
	Json := PNodeRespon{
		Times:       times,
		Type:        "Successful",
		Version:     PNode_Version,
		ServiceName: servicename,
		ServiceID:   serviceid,
	}
	jsonData, err := json.MarshalIndent(Json, "", "    ")
	if err != nil {
		log.Printf("[Info][Control][Error] Error marshaling JSON:", err)
	}
	return string(jsonData)
}

// Node S Json
func POutputConfigJson(
	appName string,
	Node_Version string,
	appHash string,
	id int,
	node_IPV4 string,
	node_IPV6 string,
	TranspondForwar bool) string {
	// JsonMssage type
	Json := PnodeResolve{
		Type:                 "PNode",
		Version:              Node_Version,
		AppName:              appName,
		AppHash:              appHash,
		ID:                   id,
		NodeIPV4:             node_IPV4,
		NodeIPV6:             node_IPV6,
		TranspondForward:     false,
		TranspondForwardPort: 1145145,
	}

	jsonData, err := json.MarshalIndent(Json, "", "    ")
	if err != nil {
		log.Printf("[Info][Control][Error] Error marshaling JSON:", err)
	}
	// log.Printf(string(jsonData))   // Console log JSON
	return string(jsonData)
}

func ParesNode(Json string) (string, map[string]interface{}, bool) {
	node := make(map[string]interface{})
	err := json.Unmarshal([]byte(Json), &node)
	if err != nil {
		log.Printf("[Info][Control][ERROR][ParesNode] ", err)
	}
	switch node["Types"] {
	case "Peer":
		return node["Types"].(string), node, true
	case "TranspondForward":
		return node["Types"].(string), node, true
	default:
		return "", nil, false
	}
}

// Node Json
func OutputCgfnodeJson(appName string, appHash string, token string, id int, node_IPV4 string, node_IPV6 string, TranspondForwar bool) string {
	Json := NodeResolve{
		Type:             "Node",
		Version:          Node_Version,
		Token:            token,
		AppName:          appName,
		APPHash:          appHash,
		AppNodeID:        id,
		NodeIPV4:         node_IPV4,
		NodeIPV6:         node_IPV6,
		TranspondForward: TranspondForwar,
	}

	jsonData, err := json.MarshalIndent(Json, "", "    ")
	if err != nil {
		log.Printf("[Info][Control][Error] Error marshaling JSON:", err)
	}
	// Console log JSON
	// log.Printf(string(jsonData))
	return string(jsonData)
}
