package parser

type ErrProtocol struct {
	Message string
}

func (err ErrProtocol) Error() string {
	return "protocol error: " + err.Message
}
