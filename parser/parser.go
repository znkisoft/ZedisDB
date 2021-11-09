package parser

type ReplyType byte

const (
	SimpleString ReplyType = '+'
	Err          ReplyType = '-'
	Integer      ReplyType = ':'
	BulkString   ReplyType = '$'
	Arrays       ReplyType = '*'
)

func (t ReplyType) String() string {
	switch t {
	default:
		return "Unknown"
	case '+':
		return "SimpleString"
	case '-':
		return "Error"
	case ':':
		return "Integer"
	case '$':
		return "BulkString"
	case '*':
		return "Array"
	}
}
