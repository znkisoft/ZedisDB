package server

import (
	"fmt"
	"github.com/znkisoft/zedisDB/lib/logger"
	"github.com/znkisoft/zedisDB/lib/utils"
	"log"
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

		// debug
		// go func() {
		// 	reader := bufio.NewReader(conn)
		// 	for {
		// 		msg, err := reader.ReadString('\n')
		// 		if err != nil {
		// 			logger.CommonLog.Println("connection closed")
		// 			os.Exit(1)
		// 		}
		// 		logger.CommonLog.Printf("(incoming message): %s", msg)
		// 		conn.Write([]byte(msg))
		// 	}
		// }()

		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	conn := NewRESPConn(c)

	for {
		v, _, err := conn.ReadValue()
		if err != nil {
			log.Fatalln(err)
		}
		values := v.Array()
		fmt.Printf("%+v\n", values)
	}

}
