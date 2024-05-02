package mssages

import(
	"encoding/hex"
	"crypto/rand"
	"time"
)

type tokenW struct {
	Type       string `json:"type"`
	Version        string `json:"version"`
	Token        string `json:"version"`
}

type tokenReply struct {
	Type       string `json:"type"`
	Version        string `json:"version"`
	VerificationCode        string `json:"version"`
}

func applyTokenDeal(tk string) {
	var apply tokenW
	err := json.Unmarshal([]byte(tk), &apply)
	if err != nil {
		log.Printf("[Info][Control][ERROR][applyTokenDeal] ",err)
	}
	if apply.Type == "token" {
		switch apply.Token{
		case 
		default nil:
			return
		}
	}
}


// Node apply Token
func applyToken(Token string) (string) {
	t := tokenW {
		Type:		"TokenApply",
		Version:           Node_Version,
		Token:           Token,
	}
	jsonData, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		log.Printf("[Info][Control][Error] Error marshaling JSON:", err)
	}
	return string(t)
}
// Node(Server) replyToken
func replyToken() (tokenReply string, token string) {
	t := makeToekn()
	tr := tokenReply {
		Type:		"TokenApply",
		Version:           Node_Version,
		VerificationCode:           t,
	}
	jsonData, err := json.MarshalIndent(tr, "", "    ")
	if err != nil {
		log.Printf("[Info][Control][Error] Error marshaling JSON:", err)
	}
	return string(tr), t
}

func makeToekn() (str string) {
	b := make([]byte, 24)
	rand.Read(b)
	str := hex.EncodeToString(b)
	return str
}
