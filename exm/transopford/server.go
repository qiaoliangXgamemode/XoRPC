package main
import (
        "fmt"
        "io"
        "net"
        "os"
		"github.com/xtaci/smux"
		"golang.org/x/net/context"
		"time"
)

type test struct {
	ip net.Conn
	pipeGet   chan *smux.Session
	ctx    context.Context
}


func main() {
	// Linux 0.0.0.0 , Windows: 127.0.0.1
	// host := "127.0.0.1"
	// port := "4001"
	l, err := net.Listen("tcp", "127.0.0.1:4001")
	if err != nil {
		fmt.Println(err, err.Error())
		os.Exit(0)
	}
	fmt.Println("[Server] Listen :4001. \r\n")
	for {
		//等待连接
		s_conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go handleConn(s_conn)
	}
}

func (c *test) getPipe() *smux.Session {
	select {
	case p := <-c.pipeGet:
		return p
	}
}

func handleConn(conn net.Conn) {

	body, _ := readMsg(conn)
	smuxConfig := smux.DefaultConfig()
	sess, err := smux.Server(conn, smuxConfig)
	if err != nil {
		fmt.Println("Error connecting to remote:", err)
	}
	stream, err := sess.AcceptStream()
	if err != nil {
		fmt.Println("Error connecting to remote:", err)
	}
	time.Sleep(time.Second * 3)
	// fmt.Printf("body:%s", string(body[:4]))
	fmt.Printf("body:%s \r\n", string(body[4:]))
	lis := ServerAddListen(string(body[4:]))

	go func(){
		conns, err := lis.Accept()
		if err == nil {
			return
		}
		ctl := &test{
			pipeGet:        make(chan *smux.Session),
			ip: conn,
		}
		go porxyConn(conns, stream, ctl)
	}()
}

func readMsg(r net.Conn) ([]byte, error) {
	length := 22
	body := make([]byte, length + 4)
	io.ReadFull(r, body)

	return body, nil
}

func porxyConn(LocalAddres net.Conn, userConn net.Conn, x *test) {
	
	p := x.getPipe()
	if p == nil {
		return
	}
	stream, err := p.OpenStream()
	if err != nil {
		return
	}
	defer stream.Close()
	go handleConnection(stream, x.ip)
}


// Add Listen
func ServerAddListen(LocalAddres string) (net.Listener) {
	fmt.Println("\r\n [Server] Listen Client:", LocalAddres)
	ip := string(net.ParseIP(LocalAddres))
	lis, err := net.Listen("tcp", ip)
	if err != nil {
		fmt.Println("\r\n Listen err: ",err)
	}
	return lis
}



func handleConnection(s_conn net.Conn, Xip net.Conn) {
	p1die := make(chan struct{})
			// 创建流通道 2
	p2die := make(chan struct{})

	go func() {
		copyData(s_conn, Xip)
		close(p1die)
	}()
	//将 反代理的端口 发送 连接方
	go func(){
		copyData(Xip, s_conn)
		close(p2die)
	}()
	select {
	case <-p1die:
	case <-p2die:
	}
}

func copyData(dst io.WriteCloser, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("Error copying data:", err)
	}
	defer dst.Close()
}