package parser

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

type RESPReader struct {
	rd *bufio.Reader
}

func NewRESPReader(rd io.Reader) *RESPReader {
	return &RESPReader{rd: bufio.NewReader(rd)}
}

func (r *RESPReader) ReadValue() (Value, int, error) {
	// handle bulk strings
	return r.readValue()
}

func (r *RESPReader) readValue() (Value, int, error) {
	t, err := r.rd.ReadByte()
	if err != nil {
		return NullValue(), 0, err
	}
	typ := ReplyType(t)

	switch typ {
	case Array:
		return r.readArrayValue()
	case SimpleString, Err:
		return r.readSimpleValue(typ)
	case Integer:
		return r.readIntegerValue()
	default:
		return NullValue(), 0, errors.New("unknown type")
	}
}

func (r *RESPReader) readArrayValue() (Value, int, error) {
	// first line is the length
	length, rn, err := r.readInt()
	n := rn
	if err != nil {
		if _, ok := err.(*ErrProtocol); ok {
			return NullValue(), n, &ErrProtocol{"invalid array length"}
		}
		return NullValue(), n, err
	}
	values := make([]Value, length)
	for i := 0; i < length; i++ {
		v, rn, err := r.readValue()
		n += rn
		if err != nil {
			return NullValue(), n, err
		}
		values[i] = v
	}
	return Value{typ: Array, array: values}, n, nil
}

func (r *RESPReader) readSimpleValue(t ReplyType) (Value, int, error) {
	line, n, err := r.readLine()
	if err != nil {
		return NullValue(), 0, err
	}
	return Value{typ: t, str: line}, n, nil
}

func (r *RESPWriter) readBulkValue() {

}

func (r *RESPReader) readIntegerValue() (Value, int, error) {
	num, n, err := r.readInt()
	if err != nil {
		return NullValue(), 0, err
	}
	return Value{typ: Integer, integer: num}, n, nil
}

func (r *RESPReader) readLine() ([]byte, int, error) {
	var (
		line []byte
		n    int
	)
	for {
		bytes, err := r.rd.ReadBytes('\n')
		if err != nil {
			return nil, 0, err
		}
		n += len(bytes)
		line := append(line, bytes...)
		if n >= 2 && line[len(line)-2] == '\r' {
			break
		}
	}
	return line[:len(line)-2], n, nil
}

func (r *RESPReader) readInt() (int, int, error) {
	line, n, err := r.readLine()
	if err != nil {
		return 0, 0, err
	}
	num, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, n, err
	}

	return int(num), n, nil
}
