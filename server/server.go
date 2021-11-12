package server

import (
	"github.com/znkisoft/zedisDB/handler"
	"github.com/znkisoft/zedisDB/lib/logger"
	"github.com/znkisoft/zedisDB/lib/utils"
	"github.com/znkisoft/zedisDB/parser"
	"io"
	"strings"
	"time"

	"net"
)

func ListenAndServe(addr string) {
	// resolve tcp addr
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	utils.CheckError(err)

	// bind listener addr
	l, err := net.ListenTCP("tcp", tcpAddr)
	utils.CheckError(err)
	defer l.Close()

	logger.CommonLog.Printf("(connected) ZedisDB is bounding to %s", addr)

	for {
		conn, err := l.AcceptTCP()
		utils.CheckError(err)

		// set timeout
		conn.SetDeadline(time.Now().Add(time.Minute))
		utils.CheckError(err)

		// err = conn.SetKeepAlivePeriod(time.Minute)
		// utils.CheckError

		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	conn := parser.NewRESPConn(c)
	defer conn.Conn.Close()

	for {
		v, _, err := conn.ReadValue()
		if err != nil {
			if err == io.EOF {
				logger.CommonLog.Printf("(disconnected) %s", c.RemoteAddr())
				return
			}
		}
		values := v.Array()
		if len(values) < 1 {
			continue
		}
		command := strings.ToUpper(values[0].String())

		if cmd, ok := handler.Router[command]; ok {
			cmd.Func(conn, values)
		} else {
			conn.Conn.Write([]byte("-ERR Unknown Command\r\n"))
		}
	}

}
