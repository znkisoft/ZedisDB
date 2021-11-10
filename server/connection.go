package server

import (
	"github.com/znkisoft/zedisDB/parser"
	"net"
)

type RESPConn struct {
	*parser.RESPReader
	*parser.RESPWriter
	net.Conn
}

func NewRESPConn(c net.Conn) *RESPConn {
	return &RESPConn{
		RESPWriter: parser.NewRESPWriter(c),
		RESPReader: parser.NewRESPReader(c),
	}
}
