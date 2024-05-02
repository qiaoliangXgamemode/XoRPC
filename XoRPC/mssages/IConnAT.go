package mssages

const (
	// Send and Recv
	TYPE_Start = 0x16
	TYPE_ENG = 0x17
)

// 1472 MUT, AES32 has 32byte. need 32 size
// to AES token.
func WriteAES256(APPHash string, conn *net.UDPConn) (string, error) {
	MUT := make([]byte, 33)
	if APPHash == nil {
		MUT[32:] := RandAes256(32)
	}
	MUT[33] := TYPE_Start
	conn.Write([]byte(MUT))
	return vrtAes256(conn)
}

func vrtAes256(conn *net.UDPConn) (string, error) {
	MUT := make([]byte, 33)
	n, remoteAddr, err := conn.ReadFromUDP(MUT)
	if err != nil {
		log.Fatal(err)
	}
	// end of trans
	if len(MUT) <= 33 && MUT[33] == 0x17 {
		return MUT[:32], nil
	}
	return _, Failed_AES
}

func remoteMoTOAES256(conn *net.UDPConn) (string, error) {
	MUT := make([]byte, 33)
	n, remoteAddr, err := conn.ReadFromUDP(MUT)
	if err != nil {
		log.Fatal(err)
	}
	// end of trans
	if len(MUT) <= 33 && MUT[33] == 0x16 {
		MUT[32:] := RandAes256(32)
		MUT[33] := TYPE_ENG
		conn.WriteToUDP([]byte(MUT), remoteAddr)
		return MUT[:32], nil
	}
	return _, Failed_AES
}

// Server recv node token deal with Apply
func NodeApplytoken(tk string, conn *net.UDPConn) {
	
}

// read node date, apply token replay.
func NodeReplytoken() {
	tk := make([]byte, 1472)
	n, err := conn.Read(tk)
	if err != nil {
		log.Fatal(err)
	}
	tk[:n]
}