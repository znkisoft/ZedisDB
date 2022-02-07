package parser

import (
	"fmt"
)

type ErrType uint32

const (
	Internal ErrType = iota + 1
	Syntax
	Protocol
	Client
)

func (t ErrType) String() string {
	switch t {
	case Internal:
		return "internal"
	case Syntax:
		return "syntax"
	case Protocol:
		return "protocol"
	case Client:
		return "client"
	}
	return ""
}

type Err struct {
	Type    ErrType
	Message string
}

func (err Err) Error() string {
	return fmt.Sprintf("-ERR [%s]: %s\r\n", err.Type.String(), err.Message)
}
