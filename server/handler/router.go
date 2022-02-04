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
	"PING":  DefaultFunc(PingCmdFunc, 1),
	"ECHO":  DefaultFunc(EchoCmdFunc, 2),
	"GET":   DefaultFunc(GetCmdFunc, 2),
	"SET":   DefaultFunc(SetCmdFunc, 3),
	"SETNX": DefaultFunc(SetNxCmdFunc, 3),
	"SETEX": DefaultFunc(SetExCmdFunc, 4),
}
