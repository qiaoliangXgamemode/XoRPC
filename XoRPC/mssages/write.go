package mssages

import ( 
	"net"
	"XoRPC/XoRPC/src/encrypt"
)
var(
	Failed_AES = errors.New("[Warning] Failed to get AES256.")
)

const (
	// Send and Recv
	TYPE_Start = 0x16
	TYPE_ENG = 0x17
)

func (*ServerMsgPool) WriteToUDPAES256(b []byte) {
	remoteAddr := ServerMsgPool.rAddr
	AES256 := ServerMsgPool.AES256
	encrypted, _ := encrypt.AES256Encrypt([]byte(b), []byte(AES256))
	ServerMsgPool.conn.WriteToUDP([]byte(encrypted), remoteAddr)
}



func Write(w io.Write, buf []byte) {

}