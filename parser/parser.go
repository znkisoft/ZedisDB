package parser

import (
	"errors"
	"strconv"
)

type ReplyType byte

type Value struct {
	typ     ReplyType
	array   []Value
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

func StringValue(s string) Value {
	return Value{typ: BulkString, str: []byte(s)}
}
func SimpleStringValue(s string) Value {
	return Value{typ: SimpleString, str: []byte(formatOneLine(s))}
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
func IntegerValue(n int) Value {
	return Value{integer: n, typ: Integer}
}

func (v Value) Bytes() []byte {
	switch v.typ {
	case SimpleString, Err, BulkString:
		return v.str
	default:
		return []byte(v.String())
	}
}
func BytesValue(b []byte) Value {
	return Value{str: b, typ: BulkString}
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
func FloatValue(f float64) Value {
	// The special precision -1 uses the smallest number of
	// digits necessary such that ParseFloat will return f exactly.
	return StringValue(strconv.FormatFloat(f, 'f', -1, 64))
}

func (v Value) IsNull() bool {
	return v.null
}
func NullValue() Value {
	return Value{typ: BulkString, null: true}
}

func (v Value) Error() error {
	switch v.typ {
	case Err:
		return errors.New(string(v.str))
	}
	return nil
}
func ErrorValue(err error) Value {
	return Value{typ: Err, str: []byte(err.Error())}
}

func (v Value) Array() []Value {
	if v.typ == Array && v.null == false {
		return v.array
	}
	return nil
}

func ArrayValue(vals []Value) Value {
	return Value{
		typ:   Array,
		array: vals,
	}
}

func (v Value) Type() ReplyType {
	return v.typ
}

// TODO
func formatOneLine(s string) string {
	return s
}
