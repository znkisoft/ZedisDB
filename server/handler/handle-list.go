package handler

import (
	"github.com/znkisoft/zedisDB/database/datastruct"
	"github.com/znkisoft/zedisDB/parser"
)

func LpushCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	// look up key in db
	_, _, err := resolveCmdArgs(cmdArgs, 2)
	if err != nil {
		return err
	}
	// if key not found, create new list

	// AMEND: optimize if storages are small  dblist -> ziplist
	return nil
}

func RpushCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
	panic("implement me")
}

func push(subject, value *datastruct.ZedisObject) error {
	panic("implement me")

}

func LpushxCmdFunc(con *parser.RESPConn, cmdArgs []parser.Value) error {
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
