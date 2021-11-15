package parser

import (
	"io"
)

type RESPWriter struct {
	R io.Writer
}

func NewRESPWriter(w io.Writer) *RESPWriter {
	return &RESPWriter{
		R: w,
	}
}

func (r *RESPWriter) WriteValue(v Value) error {
	b, err := v.MarshalRESP()
	if err != nil {
		return err
	}
	_, err = r.R.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (r *RESPWriter) WriteSimpleString(s string) error {
	return r.WriteValue(SimpleStringValue(s))
}

func (r *RESPWriter) WriteError(e ErrProtocol) error {
	// TODO rewrite inside WriteValue
	_, err := r.R.Write([]byte("-" + e.Type.String() + " " + e.Message + "\r\n"))
	return err
}

func (r *RESPWriter) WriteNull() error {
	return r.WriteValue(NullValue())
}

func (r *RESPWriter) WriteInteger(num int) error {
	return r.WriteValue(IntegerValue(num))
}

func (r *RESPWriter) WriteArray(vals []Value) error {
	return r.WriteValue(ArrayValue(vals))
}

func (r *RESPWriter) WriteBulkString(s string) error {
	return r.WriteValue(StringValue(s))
}
