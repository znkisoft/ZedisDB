package handler

import (
	"github.com/znkisoft/zedisDB/parser"
)

type Cmd struct {
	Name     string
	Arity    int
	Func     CmdFunc
	Category CmdType
}

type CmdFunc func(con *parser.RESPConn, cmdArgs []parser.Value)

type CmdType int

const (
	Server = iota
	Connection
	Strings
	Lists
	Hashes
	Keys
)

var Router = map[string]Cmd{
	"PING": {Name: "ping", Arity: 0, Func: DefaultFunc(PingCmdFunc, 1), Category: Strings},
	"ECHO": {Name: "echo", Arity: 1, Func: DefaultFunc(EchoCmdFunc, 2), Category: Strings},
	// "GET": {Name: "get", Arity: 2, Func: CmdFuncGet, Category: Strings},
}
