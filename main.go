package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	var (
		port string
	)
	flag.StringVar(&port, "port", "6379", "define the port(e.g. 6379)")
	flag.Parse()

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%s", port))
	CheckError(err)

	listener, err := net.ListenTCP("tcp", addr)
	CheckError(err)
	defer listener.Close()

	log.Printf("[connected] bound to %q", listener.Addr())

	for {
		tcpConn, err := listener.AcceptTCP()
		CheckError(err)

		// err = tcpConn.SetKeepAlive(true)
		// CheckError(err)

		// err = tcpConn.SetKeepAlivePeriod(time.Minute)
		// CheckError(err)

		go func(c net.Conn) {
			defer c.Close()
			buf := make([]byte, 1<<10) // 1KB
			for {
				// there’s no guarantee that all these bytes will arrive at the same time
				n, err := c.Read(buf)
				CheckError(err)

				handle(c, n, buf[:n])
			}
		}(tcpConn)

	}
}

func handle(conn net.Conn, length int, data []byte) {
	var decoder Decoder
	msg := NewMessage(data, length)

	msgType, err := decoder.CheckType(msg)
	CheckError(err)

	switch msgType {
	case RedisReplyString:
		conn.Write(MsgPong)
	case RedisReplyArray:
		conn.Write(MsgPong)
	default:
	}
}
