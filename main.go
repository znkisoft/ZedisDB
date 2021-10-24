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

	log.Printf("[running] bound to %q", listener.Addr())

	for {
		tcpConn, err := listener.AcceptTCP()
		CheckError(err)
		go func(c net.Conn) {
			buf := make([]byte, 100)
			length, err := c.Read(buf)
			CheckError(err)

			handle(c, length, buf)
		}(tcpConn)

	}
}

func handle(conn net.Conn, length int, msg []byte) {
	fmt.Printf("Received data: %v, length: %d\n", string(msg[:length]), length)
	// conn.Write([]byte("+PONG\n"))
	conn.Close()
}
