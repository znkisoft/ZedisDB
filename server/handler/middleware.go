package handler

import (
	"strconv"

	"github.com/znkisoft/zedisDB/logger"
	"github.com/znkisoft/zedisDB/parser"
)

func ArgsCheckFunc(fn CmdFunc, argLimit int) CmdFunc {
	return func(con *parser.RESPConn, cmdArgs []parser.Value) error {
		logger.CommonLog.Printf("(handler)[%s] %s", con.Conn.RemoteAddr(), cmdArgs)
		if len(cmdArgs) != argLimit {
			return parser.Err{Type: parser.Client, Message: "wrong number of arguments, expect " + strconv.Itoa(argLimit)}
		}
		return fn(con, cmdArgs)
	}
}
