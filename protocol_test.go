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
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
