package server

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"

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
		// utils.CheckError
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			// close connection if end with io.EOF
			if err == io.EOF {
				log.Println("connection closed")
				os.Exit(1)
			} else {
				log.Println(err)
			}
			return
		}

		// TODO resolve comming request with payload
		bytes := []byte(msg)
		conn.Write(bytes)
	}
}
