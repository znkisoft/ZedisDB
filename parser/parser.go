package parser

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"strconv"
)

type RESPConn struct {
	*RESPReader
	*RESPWriter
	Conn net.Conn
}

func NewRESPConn(c net.Conn) *RESPConn {
	return &RESPConn{
		RESPWriter: NewRESPWriter(c),
		RESPReader: NewRESPReader(c),
		Conn:       c,
	}
}

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
	ErrStr       ReplyType = '-'
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
	case SimpleString, ErrStr:
		return string(v.str)
	case Integer:
		return strconv.Itoa(v.integer)
	case BulkString:
		return string(v.str)
	case Array:
		return fmt.Sprintf("%v", v.array)
	}
	return ""
}

func StringValue(s string) Value {
	return Value{typ: BulkString, str: []byte(s)}
}

func SimpleStringValue(s string) Value {
	return Value{typ: SimpleString, str: []byte(formatOneLine(s))}
}

func EmptyStringValue() Value {
	return Value{typ: BulkString, str: []byte("0\r\n")}
}

func NullStringValue() Value {
	return Value{typ: BulkString, null: true, str: []byte("-1")}
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
	case SimpleString, ErrStr, BulkString:
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

// BoolValue returns a RESP integer representation of a bool.
func BoolValue(t bool) Value {
	if t {
		return Value{typ: ':', integer: 1}
	}
	return Value{typ: ':', integer: 0}
}

func (v Value) Bool() bool {
	return v.Integer() != 0
}

func (v Value) IsNull() bool {
	return v.null
}

// NullValue returns a null value, origin: "$-1\r\n"
func NullValue() Value {
	return Value{typ: BulkString, null: true}
}

// MultiBulkValue returns a RESP array which contains one or more bulk strings. l.
func MultiBulkValue(commandName string, args ...interface{}) Value {
	vals := make([]Value, len(args)+1)
	vals[0] = StringValue(commandName)
	for i, arg := range args {
		if rval, ok := arg.(Value); ok && rval.Type() == BulkString {
			vals[i+1] = rval
			continue
		}
		switch arg := arg.(type) {
		default:
			vals[i+1] = StringValue(fmt.Sprintf("%v", arg))
		case []byte:
			vals[i+1] = StringValue(string(arg))
		case string:
			vals[i+1] = StringValue(arg)
		case nil:
			vals[i+1] = NullValue()
		}
	}
	return ArrayValue(vals)
}

func (v Value) Error() error {
	switch v.typ {
	case ErrStr:
		return errors.New(string(v.str))
	}
	return nil
}
func ErrorValue(err error) Value {
	return Value{typ: ErrStr, str: []byte(err.Error())}
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

// AnyValue returns a RESP value from an interface. This function infers the types. Arrays are not allowed.
func AnyValue(v interface{}) Value {
	switch v := v.(type) {
	default:
		return StringValue(fmt.Sprintf("%v", v))
	case nil:
		return NullValue()
	case int:
		return IntegerValue(v)
	case uint:
		return IntegerValue(int(v))
	case int8:
		return IntegerValue(int(v))
	case uint8:
		return IntegerValue(int(v))
	case int16:
		return IntegerValue(int(v))
	case uint16:
		return IntegerValue(int(v))
	case int32:
		return IntegerValue(int(v))
	case uint32:
		return IntegerValue(int(v))
	case int64:
		return IntegerValue(int(v))
	case uint64:
		return IntegerValue(int(v))
	case bool:
		return BoolValue(v)
	case float32:
		return FloatValue(float64(v))
	case float64:
		return FloatValue(float64(v))
	case []byte:
		return BytesValue(v)
	case string:
		return StringValue(v)
	}
}

func (v Value) Type() ReplyType {
	return v.typ
}

// Equals compares one value to another value.
func (v Value) Equals(value Value) bool {
	data1, err := v.MarshalRESP()
	if err != nil {
		return false
	}
	data2, err := value.MarshalRESP()
	if err != nil {
		return false
	}
	return string(data1) == string(data2)
}

func formatOneLine(s string) string {
	var buf bytes.Buffer
	for _, c := range s {
		switch c {
		case '\r', '\n':
			buf.WriteString("\\r\\n")
		default:
			buf.WriteRune(c)
		}
	}
	return buf.String()
}

func (v Value) MarshalRESP() ([]byte, error) {
	return marshalAnyRESP(v)
}

func marshalAnyRESP(v Value) ([]byte, error) {
	switch v.typ {
	case SimpleString, ErrStr:
		return marshalSimpleRESP(v.typ, v.str)
	case Integer:
		return marshalSimpleRESP(v.typ, []byte(strconv.Itoa(v.integer)))
	case BulkString:
		return marshalBulkRESP(v)
	case Array:
		return marshalArrayRESP(v)
	default:
		if v.typ == 0 && v.null {
			return []byte("$-1\r\n"), nil
		}
		return nil, errors.New("unknown resp type")
	}
}

func marshalSimpleRESP(typ ReplyType, b []byte) ([]byte, error) {
	return []byte(fmt.Sprintf("%c%s\r\n", typ, b)), nil
}

func marshalBulkRESP(v Value) ([]byte, error) {
	if v.IsNull() {
		return []byte("$-1\r\n"), nil
	}
	szb := []byte(strconv.FormatInt(int64(len(v.str)), 10))
	bb := make([]byte, 5+len(szb)+len(v.str))
	bb[0] = '$'
	copy(bb[1:], szb)
	bb[1+len(szb)+0] = '\r'
	bb[1+len(szb)+1] = '\n'
	copy(bb[1+len(szb)+2:], v.str)
	bb[1+len(szb)+2+len(v.str)+0] = '\r'
	bb[1+len(szb)+2+len(v.str)+1] = '\n'
	return bb, nil
}

func marshalArrayRESP(v Value) ([]byte, error) {
	if v.IsNull() {
		return []byte("*-1\r\n"), nil
	}
	var buf bytes.Buffer
	szb := []byte(strconv.FormatInt(int64(len(v.array)), 10))
	buf.Grow(3 + len(szb) + 16*len(v.array))
	buf.WriteByte('*')
	buf.Write(szb)
	buf.WriteByte('\r')
	buf.WriteByte('\n')
	for i := 0; i < len(v.array); i++ {
		data, err := v.array[i].MarshalRESP()
		if err != nil {
			return nil, err
		}
		buf.Write(data)
	}
	return buf.Bytes(), nil
}
