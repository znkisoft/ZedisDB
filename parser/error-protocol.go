package parser

// const (
// 	?Common = iota
// 	Syntax
// 	Protocol
// 	Param
// 	Server
// )

type ErrProtocol struct {
	Message string
}

func (err ErrProtocol) Error() string {
	return "protocol error: " + err.Message
}
