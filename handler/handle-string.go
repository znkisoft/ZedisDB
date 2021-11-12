package handler

import (
	"github.com/znkisoft/zedisDB/parser"
)

func PingCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) {
	con.WriteSimpleString("PONG")
}

func EchoCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) {
	con.WriteSimpleString(cmdArgs[1].String())
}
