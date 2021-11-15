package parser

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
		return "Server"
	case Syntax:
		return "Syntax"
	case Protocol:
		return "Protocol"
	case Param:
		return "Param"
	}
	return ""
}

type ErrProtocol struct {
	Type    ErrType
	Message string
}

func (err ErrProtocol) Error() string {
	return "protocol error: " + err.Message
}
