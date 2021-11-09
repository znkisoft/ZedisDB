package parser

import (
	"errors"
	"strconv"
)

type ReplyType byte

type Value struct {
	typ     ReplyType
	data    []Value
	null    bool
	str     []byte
	integer int
}

const (
	SimpleString ReplyType = '+'
	Err          ReplyType = '-'
	Integer      ReplyType = ':'
	BulkString   ReplyType = '$'
	Array        ReplyType = '*'
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

func (v Value) String() string {
	switch v.typ {
	case SimpleString, Err:
		return string(v.str)
	case Integer:
		return strconv.Itoa(v.integer)
	case BulkString:
		// TODO array parsing
		return ""
	}
	return ""
}

func (v Value) Integer() int {
	switch v.typ {
	case Integer:
		return v.integer
	default:
		n, _ := strconv.Atoi(v.String())
		return n
	}
}

func (v Value) Bytes() []byte {
	switch v.typ {
	case SimpleString, Err, BulkString:
		return v.str
	default:
		return []byte(v.String())
	}
}

func (v Value) Float() float64 {
	switch v.typ {
	case Integer:
		return float64(v.integer)
	default:
		n, _ := strconv.ParseFloat(v.String(), 64)
		return n
	}
}

func (v Value) IsNull() bool {
	return v.null
}

func (v Value) Error() error {
	switch v.typ {
	case Err:
		return errors.New(string(v.str))
	}
	return nil
}
