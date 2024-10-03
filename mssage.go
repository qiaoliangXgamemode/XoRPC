package XoRPC

import (
	"io"
	"time"

	"github.com/quic-go/quic-go"
)

// Identification code
// FindNodeCode | 201 | find node on, | 200 | find node off.
var (
	FINDNODECODE_OFF = uint8(100)
	FINDNODECODE_ON  = uint8(101)
	FINDLTT          = uint8(6)
)

type Pong struct {
}

func newRPChead(FINDNODECODE uint8, FINDLTT uint8) []byte {
	head := make([]byte, 2)
	head[0] = FINDNODECODE_OFF
	head[1] = FINDLTT
	return head
}
func newFindNodeMssagePack(p []byte) []byte {
	var length int
	body := p
	length = len(body)
	x := make([]byte, length+2)
	c := newRPChead(FINDNODECODE_OFF, 6)
	copy(x[:2], c)
	copy(x[2:], body)

	return x
}
func WriteStreamQuic(stream quic.Stream, p []byte) []byte {
	// stream.SetWriteDeadline(time.Now().Add(time.Second * 12))
	if _, ok := stream.Write(p); ok == io.EOF {

	}
	return ReadStreamQuic(stream)
}

func ReadStreamQuic(stream quic.Stream) []byte {
	stream.SetReadDeadline(time.Now().Add(time.Second))
	buf := make([]byte, 1024)
	n, _ := io.Reader.Read(stream, buf)
	// r := buf[:n]
	return buf[:n]
}

func ReadRPCcodeStreamQuic(stream quic.Stream) ([]byte, int) {
	buf := make([]byte, 1024)
	n, _ := io.Reader.Read(stream, buf)

	r := buf[:n]
	if r[0] == FINDNODECODE_OFF {
		return r[2:], 201
	}
	// if r := buf[:n]; string(r[:1]) == "201" {
	// 	return buf[:n+2], 201
	// }
	// r := buf[:n]
	// slog.Debug("Debug message", "mssage")
	// return r, 200
	// r := buf[:n]
	return buf[:n+2], 200
}
