package handler

import (
	"github.com/znkisoft/zedisDB/parser"
)

func LpushCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	if len(cmdArgs) < 2 {
		return parser.Err{Type: parser.Client, Message: "wrong number of arguments for 'lpush' command"}
	}

	return nil
}

func LpushxCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	panic("implement me")
}

func RpushCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	panic("implement me")
}

func RpushxCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	panic("implement me")
}

func LpopCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	panic("implement me")
}

func RpopCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	panic("implement me")
}
func RpoplpushCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	panic("implement me")
}
func LremCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	panic("implement me")
}
func LlenCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	panic("implement me")
}
func LindexCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	panic("implement me")
}
func LsetCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	panic("implement me")
}
func LrangeCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	panic("implement me")
}
