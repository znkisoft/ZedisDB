package main

import (
	"flag"

	"github.com/znkisoft/zedisDB/server"
)

func main() {
	var (
		port string
	)
	flag.StringVar(&port, "port", "6379", "define the port(e.g. 6379)")
	flag.Parse()

	server.ListenAndServe("127.0.0.1:" + port)
}

// func handle(conn net.Conn, length int, data []byte) {
// 	var (
// 		decoder  RESPDecode
// 		command  Command
// 		response []byte
// 	)
// 	msg := NewMessage(data, length)
// 	msgType := decoder.CheckType(msg)

// 	switch msgType {
// 	case ZedisReplyString:
// 		str, _ := decoder.DecodeSimpleString(msg)
// 		response = command.HandleCommand(str, nil)
// 		conn.Write(response)
// 	case ZedisReplyArray:
// 		str, _ := decoder.DecodeArray(msg)
// 		response = command.HandleCommand(str[0], str[1:])
// 		conn.Write(response)
// 	case ZedisReplyUnknown:
// 		conn.Write([]byte("*1unknown\r\ncommand\r\n"))
// 	default:
// 	}
// }
