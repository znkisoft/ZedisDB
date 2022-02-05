package handler

import (
	"strconv"

	"github.com/znkisoft/zedisDB/database/container"

	"github.com/znkisoft/zedisDB/parser"
	"github.com/znkisoft/zedisDB/pkg/logger"
)

func DefaultFunc(fn CmdFunc, argLimit int) CmdFunc {
	return func(con *parser.RESPConn, cmdArgs []parser.Value) error {
		logger.CommonLog.Printf("(handler)[%s] %s", con.Conn.RemoteAddr(), cmdArgs)
		if len(cmdArgs) != argLimit {
			return parser.ErrProtocol{Type: parser.Client, Message: "wrong number of arguments, expect " + strconv.Itoa(argLimit)}
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
	val, ok := db.Get(key)
	logger.CommonLog.Printf("[GET](%t): %s", ok, key)
	if !ok {
		return con.WriteNull()
	}
	return con.WriteValue(parser.AnyValue(val.Ptr()))
}

/*SetCmdFunc
pattern: SET key value [NX] [XX] [EX <seconds>] [PX <milliseconds>]
todo: nx xx ex px
*/
func SetCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	key := cmdArgs[1].String()
	value := cmdArgs[2]
	o := container.CreateZedisObject(container.StringTyp, value.String())
	err := db.Set(key, o)
	logger.CommonLog.Printf("[SET]key: %s, value: %s", key, value)
	if err != nil {
		return parser.ErrProtocol{Type: parser.Internal, Message: err.Error()}
	}
	return con.WriteSimpleString("OK")
}

func SetNxCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	key := cmdArgs[1].String()
	if _, found := db.Get(key); found {
		return con.WriteInteger(0)
	}
	o := container.CreateZedisObject(container.StringTyp, cmdArgs[2].String())
	err := db.Set(key, o)
	logger.CommonLog.Printf("[SETNX]key: %s, value: %s", key, cmdArgs[2])
	if err != nil {
		return parser.ErrProtocol{Type: parser.Internal, Message: err.Error()}
	}
	return con.WriteInteger(1)
}

func SetExCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	return nil
}

// TODO
// func setGenericCommand(con *parser.RESPConn, cmdArgs []parser.Value) error {
// }
