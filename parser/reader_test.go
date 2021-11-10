package parser

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestRESPReader_ReadValue(t *testing.T) {
	type fields struct {
		rd *bufio.Reader
	}
	tests := []struct {
		name      string
		fields    fields
		wantValue Value
		wantN     int
		wantErr   bool
	}{
		{
			name: "simple string",
			fields: fields{
				rd: bufio.NewReader(bufio.NewReader(strings.NewReader("+OK\r\n"))),
			},
			wantValue: Value{
				typ: SimpleString,
				str: []byte("OK"),
			},
			wantN:   4,
			wantErr: false,
		},
		{
			name: "integer",
			fields: fields{
				rd: bufio.NewReader(bufio.NewReader(strings.NewReader(":123\r\n"))),
			},
			wantValue: Value{
				typ:     Integer,
				integer: 123,
			},
			wantN:   5,
			wantErr: false,
		},
		{
			name: "bulk string",
			fields: fields{
				rd: bufio.NewReader(bufio.NewReader(strings.NewReader("*2\r\n$3\r\nGET\r\n$4\r\nUSER\r\n\r\n"))),
			},
			wantValue: Value{
				typ: Array,
				array: []Value{
					{
						typ: SimpleString,
						str: []byte("GET"),
					},
					{
						typ: SimpleString,
						str: []byte("USER"),
					},
				},
			},
			wantN:   20,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RESPReader{
				rd: tt.fields.rd,
			}
			got, got1, err := r.ReadValue()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantValue) {
				t.Errorf("ReadValue() got = %v, wantValue %v", got, tt.wantValue)
			}
			if got1 != tt.wantN {
				t.Errorf("ReadValue() got1 = %v, wantValue %v", got1, tt.wantN)
			}
		})
	}
}