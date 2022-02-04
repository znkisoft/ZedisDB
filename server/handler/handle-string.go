package handler

import (
	"strconv"

	"github.com/znkisoft/zedisDB/lib/logger"
	"github.com/znkisoft/zedisDB/parser"
)

func DefaultFunc(fn CmdFunc, argLimit int) CmdFunc {
	return func(con *parser.RESPConn, cmdArgs []parser.Value) error {
		logger.CommonLog.Printf("(handler)[%s] %s", con.Conn.RemoteAddr(), cmdArgs)
		if len(cmdArgs) != argLimit {
			return con.WriteError(parser.ErrProtocol{Type: parser.Client, Message: "wrong number of arguments, expect " + strconv.Itoa(argLimit)})
		}
		return fn(con, cmdArgs)
	}
}

func PingCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	return con.WriteValue(parser.StringValue("PONG"))
}

func EchoCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	return con.WriteSimpleString(cmdArgs[1].String())
}

func GetCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	key := cmdArgs[1].String()
	val, ok := db.Dict.Get(key)
	logger.CommonLog.Printf("[GET](%t): %s", ok, key)
	if !ok {
		return con.WriteNull()
	}
	return con.WriteValue(parser.AnyValue(val))
}

/*SetCmdFunc
pattern: SET key value [NX] [XX] [EX <seconds>] [PX <milliseconds>]
todo: nx xx ex px
*/
func SetCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	key := cmdArgs[1].String()
	value := cmdArgs[2]
	err := db.Dict.Set(key, value)
	logger.CommonLog.Printf("[SET]key: %s, value: %s", key, value)
	if err != nil {
		return con.WriteError(parser.ErrProtocol{Type: parser.Internal, Message: err.Error()})
	}
	return con.WriteSimpleString("OK")
}

func SetNxCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	key := cmdArgs[1].String()
	if _, found := db.Dict.Get(key); found {
		return con.WriteInteger(0)
	}
	err := db.Dict.Set(key, cmdArgs[2])
	logger.CommonLog.Printf("[SETNX]key: %s, value: %s", key, cmdArgs[2])
	if err != nil {
		return con.WriteError(parser.ErrProtocol{Type: parser.Internal, Message: err.Error()})
	}
	return con.WriteInteger(1)
}

func SetExCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	return nil
}

// TODO
// func setGenericCommand(con *parser.RESPConn, cmdArgs []parser.Value) error {
// }
