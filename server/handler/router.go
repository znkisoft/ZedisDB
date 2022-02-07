package handler

import (
	"github.com/znkisoft/zedisDB/database"
	"github.com/znkisoft/zedisDB/parser"
)

// TODO not sure if this is the best way to do this
var db = database.NewDb()

type CmdFunc func(con *parser.RESPConn, cmdArgs []parser.Value) error

type CmdType int

var Router = map[string]CmdFunc{
	"PING":  ArgsCheckFunc(PingCmdFunc, 1),
	"ECHO":  ArgsCheckFunc(EchoCmdFunc, 2),
	"GET":   ArgsCheckFunc(GetCmdFunc, 2),
	"SET":   ArgsCheckFunc(SetCmdFunc, 3),
	"SETNX": ArgsCheckFunc(SetNxCmdFunc, 3),
	"SETEX": ArgsCheckFunc(SetExCmdFunc, 4),
	"LPUSH": LpushCmdFunc,
}
