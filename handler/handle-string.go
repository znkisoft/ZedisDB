package handler

import (
	"github.com/znkisoft/zedisDB/lib/logger"
	"github.com/znkisoft/zedisDB/parser"
)

func DefaultFunc(fn CmdFunc) CmdFunc {
	return func(con *parser.RESPConn, cmdArgs []parser.Value) {
		logger.CommonLog.Printf("(handler)[%s] %s", con.Conn.RemoteAddr(), cmdArgs)
		fn(con, cmdArgs)
	}
}

func PingCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) {
	con.WriteSimpleString("PONG")
}

func EchoCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) {
	if len(cmdArgs) != 2 {
		con.WriteError(parser.Param, "wrong number of arguments for 'echo' command")
		return
	}
	con.WriteSimpleString(cmdArgs[1].String())
}
