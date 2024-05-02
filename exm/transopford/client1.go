package main
import (
        "fmt"
        "io"
        "net"
		"time"
		"github.com/xtaci/smux"
		// "golang.org/x/net/context"
)



func main() {
	d_tcpAddr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:4001")
	remoteAddr, err := net.DialTCP("tcp", nil, d_tcpAddr)
	if err != nil {
		fmt.Println(err)
		remoteAddr.Close()
	}
	message := `127.0.0.1:4399`
	// stream.Write([]byte(message))
	length := len(message)
	x := make([]byte, length+4)
	x[0] = uint8(1)
	x[1] = uint8(4)
	x[2] = uint8(8)
	x[3] = uint8(9)
	copy(x[4:], message)
	WriteMsg(remoteAddr, x)
	// for i := 0; i <= 100; i++{
	// 	WriteMsg(stream, x)
	// }
	
	time.Sleep(time.Second * 3)

	smuxConfig := smux.DefaultConfig()
	// 设置最大的Buffer
	smuxConfig.MaxReceiveBuffer = 1194304
	session, err := smux.Client(remoteAddr, smuxConfig)
	if err != nil {
        panic(err)
    }
	stream, err := session.OpenStream()
    if err != nil {
        panic(err)
    }
	stream.SetWriteDeadline(time.Now().Add(time.Second * 12))
	// 本地反代理的服务器

	localAddr := "127.0.0.1:25565"
	// writeFull(stream, []byte(message))
	// writeFull(stream, []byte(message))
	handleConnection(stream, localAddr)
}

func WriteMsg(w net.Conn, x []byte) error {
	w.SetWriteDeadline(time.Now().Add(time.Second * 12))
	// io写入
	writeFull(w, x)

	w.SetWriteDeadline(time.Time{})
	return nil
}

// io write
func writeFull(w io.Writer, buf []byte) (err error) {
	_, err = w.Write(buf)
	if err != nil {
		fmt.Println("Error connecting to remote:", err)
	}
	return
}

func pipeHandShake(conn net.Conn) (*smux.Session, error) {
	smuxConfig := smux.DefaultConfig()
	mux, err := smux.Server(conn, smuxConfig)
	return mux, err
}

func handleConnection(remoteAddr net.Conn, LoctAddres string) {
	
	// smux
	remoteAddrs, err := pipeHandShake(remoteAddr)
	stream, err := remoteAddrs.AcceptStream()
	if err != nil {
		fmt.Println("Error connecting to remote:", err)
		return
	}
	loctconn, err := net.Dial("tcp", LoctAddres)
	if err != nil {
		fmt.Println("Error connecting to remote:", err)
		loctconn.Close()
		return
	}
	// defer remoteConn.Close()

	fmt.Printf("Connected to remote %s\n", remoteAddr.LocalAddr())
	fmt.Printf("Connected to Local Server %s\n", loctconn.LocalAddr())
	fmt.Printf("Connected to Local Server Listen: %s\n", loctconn.RemoteAddr())
	// 启动两个 goroutine 分别进行数据转发
	for{
		p1die := make(chan struct{})
		// 创建流通道 2
		p2die := make(chan struct{})
		go copyData(stream, loctconn)
		go copyData(loctconn, stream)
		select {
		case <-p1die:
		case <-p2die:
		}
	}
}

func copyData(dst io.WriteCloser, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("Error copying data:", err)
	}
	defer dst.Close()
}