package server

import (
	"bufio"
	"github.com/znkisoft/zedisDB/lib/logger"
	"github.com/znkisoft/zedisDB/lib/utils"
	"github.com/znkisoft/zedisDB/parser"
	"io"
	"os"
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
		conn.SetDeadline(time.Now().Add(time.Second * 30))
		utils.CheckError(err)

		// err = conn.SetKeepAlivePeriod(time.Minute)
		// utils.CheckError

		go func() {
			reader, writer := bufio.NewReader(conn), bufio.NewWriter(conn)
			for {

				msg, err := reader.ReadString('\n')
				if err != nil {
					if _, ok := err.(*parser.ErrProtocol); ok {
						writer.WriteString("-ERR" + err.Error() + "\r\n")
						writer.Flush()

						// close connection if end with io.EOF
					} else if err == io.EOF {
						logger.CommonLog.Println("connection closed")
						os.Exit(1)
					} else {
						writer.WriteString("-ERR unknown error\r\n")
						writer.Flush()
					}
					return
				}
				// debug
				logger.CommonLog.Printf("(incoming message): %s", msg)

				// TODO resolve coming request with payload
				// conn.Write(bytes)
				// writer.WriteString("+OK\r\n")
				writer.WriteString(msg)
				writer.Flush()
			}
		}()
	}
}
