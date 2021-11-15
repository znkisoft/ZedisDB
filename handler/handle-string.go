package handler

import (
	"github.com/znkisoft/zedisDB/lib/logger"
	"github.com/znkisoft/zedisDB/parser"
)

func DefaultFunc(fn CmdFunc, argLimit int) CmdFunc {
	return func(con *parser.RESPConn, cmdArgs []parser.Value) {
		logger.CommonLog.Printf("(handler)[%s] %s", con.Conn.RemoteAddr(), cmdArgs)
		if len(cmdArgs) > argLimit {
			con.WriteError(parser.ErrProtocol{Type: parser.Param, Message: "too many arguments"})
			return
		}
		fn(con, cmdArgs)
	}
}

func PingCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) {
	con.WriteValue(parser.StringValue("PONG"))
}

func EchoCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) {
	con.WriteSimpleString(cmdArgs[1].String())
}
