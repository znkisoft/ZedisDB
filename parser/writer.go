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
	// switch v.Type() {
	// case TypeString:
	// 	return r.WriteString(v.String())
	// case TypeError:
	// 	return r.WriteError(v.Error())
	// case TypeInteger:
	// 	return r.WriteInteger(v.Integer())
	// case TypeNil:
	// 	return r.WriteNil()
	// case TypeArray:
	// 	return r.WriteArray(v.Array())
	// case TypeBulkString:
	// 	return r.WriteBulkString(v.BulkString())
	// case TypeSimpleString:
	// 	return r.WriteSimpleString(v.SimpleString())
	// default:
	// 	return ErrUnknownType
	// }
	return nil
}

// TODO: implement
// func marshalSimpleRESP(s string) []byte {
// 	return []byte(s + "\r\n")
// }

// func marshalBulkRESP(s string) []byte {
// 	return []byte("$" + string(len(s)) + "\r\n" + s + "\r\n")
// }

// func marshalArrayRESP(s []string) []byte {
// 	var buf []byte
// 	buf = append(buf, []byte("*"+string(len(s))+"\r\n"))
// 	for _, v := range s {
// 		buf = append(buf, marshalBulkRESP(v)...)
// 	}
// 	return buf
// }

func (r *RESPWriter) WriteSimpleString(s string) {
	r.R.Write([]byte("+" + s + "\r\n"))
}
