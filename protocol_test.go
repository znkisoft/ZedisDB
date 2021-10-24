package main

import (
	"reflect"
	"testing"
)

func TestRESPDecode_DecodeArray(t *testing.T) {
	type args struct {
		msg Message
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &RESPDecode{}
			got, err := d.DecodeArray(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeArray() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRESPDecode_DecodeBulkString(t *testing.T) {
	type args struct {
		msg Message
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"case 0", args{msg: []byte("$0\r\n\r\n")}, "", false},
		{"case 1", args{msg: []byte("$2\r\nOK\r\n")}, "OK", false},
		{"case 2", args{msg: []byte("$2\r\nHE\r\n")}, "HE", false},
		{"case 3", args{msg: []byte("$3\r\nHEY\r\n")}, "HEY", false},
		{"case 4", args{msg: []byte("$10\r\nabcdefghij\r\n")}, "abcdefghij", false},
		{"case 5 (incomplete resp)", args{msg: []byte("$")}, "", true},
		{"case 6 (incomplete resp)", args{msg: []byte("$2")}, "", true},
		{"case 7 (incomplete resp)", args{msg: []byte("$2\r")}, "", true},
		{"case 8 (incomplete resp)", args{msg: []byte("$2\r\n")}, "", true},
		{"case 9 (incomplete resp)", args{msg: []byte("$2\r\nOK")}, "", true},
		{"case 10 (incomplete resp)", args{msg: []byte("$2\r\nOK\r")}, "", true},
		// TODO
		// {"case 10 (null bulk string, should return to client as a nil object )",
		// 	args{msg: []byte("$-1\r\n")},"",true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &RESPDecode{}
			got, err := d.DecodeBulkString(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeBulkString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeBulkString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRESPDecode_DecodeSimpleString(t *testing.T) {
	type args struct {
		msg Message
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"case 1", args{msg: []byte("+OK\r\n")}, "OK", false},
		{"case 2", args{msg: []byte("+HEY\r\n")}, "HEY", false},
		{"case 3 (incomplete RESP)", args{msg: []byte("+")}, "", true},
		{"case 4 (incomplete RESP)", args{msg: []byte("+OK\r")}, "", true},
		{"case 5 (incomplete RESP)", args{msg: []byte("+OK\r")}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &RESPDecode{}
			got, err := d.DecodeSimpleString(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeSimpleString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeSimpleString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
