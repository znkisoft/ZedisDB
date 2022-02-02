package parser

import (
	"fmt"
)

type ErrType uint32

const (
	Server ErrType = iota

	Syntax
	Protocol
	Param
)

func (t ErrType) String() string {
	switch t {
	case Server:
		return "server"
	case Syntax:
		return "syntax"
	case Protocol:
		return "protocol"
	case Param:
		return "param"
	}
	return ""
}

type ErrProtocol struct {
	Type    ErrType
	Message string
}

func (err ErrProtocol) Error() string {
	return fmt.Sprintf("protocol error[%s]: %s", err.Type.String(), err.Message)
}
