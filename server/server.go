package server

import (
	"net"

	"github.com/znkisoft/zedisDB/lib/logger"
	"github.com/znkisoft/zedisDB/lib/utils"
)

func ListenAndServe(addr string) {
	// resolve tcp addr
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	utils.CheckError(err)

	// bind listener addr
	listener, err := net.ListenTCP("tcp", tcpAddr)
	utils.CheckError(err)
	defer listener.Close()

	logger.CommonLog.Printf("(connected) ZedisDB is bounding to %s", addr)

	for {
		conn, err := listener.AcceptTCP()
		utils.CheckError(err)

		// err = tcpConn.SetKeepAlive(true)
		// utils.CheckError(err)

		// err = conn.SetKeepAlivePeriod(time.Minute)
		// utils.CheckError(err)

		go func(c net.Conn) {
			defer c.Close()
			buf := make([]byte, 1<<10) // 1KB
			for {
				n, err := c.Read(buf)
				utils.CheckError(err)

				handle(c, n, buf[:n])
			}
		}(conn)

	}
}

func handle(conn net.Conn, length int, data []byte) {
	conn.Write([]byte("+pong\r\n"))

	defer conn.Close()
}
