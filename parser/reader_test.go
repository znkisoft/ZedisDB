package parser

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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
				rd: bufio.NewReader(strings.NewReader("+OK\r\n")),
			},
			wantValue: Value{
				typ: SimpleString,
				str: []byte("OK"),
			},
			wantN:   5,
			wantErr: false,
		},
		{
			name: "integer",
			fields: fields{
				rd: bufio.NewReader(strings.NewReader(":123\r\n")),
			},
			wantValue: Value{
				typ:     Integer,
				integer: 123,
			},
			wantN:   6,
			wantErr: false,
		},
		{
			name: "bulk string",
			fields: fields{
				rd: bufio.NewReader(strings.NewReader("*2\r\n$3\r\nGET\r\n$4\r\nUSER\r\n\r\n")),
			},
			wantValue: Value{
				typ: Array,
				array: []Value{
					{
						typ: BulkString,
						str: []byte("GET"),
					},
					{
						typ: BulkString,
						str: []byte("USER"),
					},
				},
			},
			wantN:   23,
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
			if !assert.Equal(t, tt.wantValue, got) {
				t.Errorf("ReadValue() got = %v, wantValue %v", got, tt.wantValue)
			}
			if got1 != tt.wantN {
				t.Errorf("ReadValue() got1 = %v, wantValue %v", got1, tt.wantN)
			}
		})
	}
}

func TestIntegers(t *testing.T) {
	sum := 0
	data := []byte(":1234567\r\n:-90898\r\n:0\r\n")
	r := NewRESPReader(bytes.NewBuffer(data))
	v, rn, err := r.ReadValue()
	sum += rn
	if err != nil {
		t.Fatal(err)
	}
	if v.Integer() != 1234567 {
		t.Fatalf("invalid integer: expected %d, got %d", 1234567, v.Integer())
	}
	v, rn, err = r.ReadValue()
	sum += rn
	if err != nil {
		t.Fatal(err)
	}
	if v.Integer() != -90898 {
		t.Fatalf("invalid integer: expected %d, got %d", -90898, v.Integer())
	}
	v, rn, err = r.ReadValue()
	sum += rn
	if err != nil {
		t.Fatal(err)
	}
	if v.Integer() != 0 {
		t.Fatalf("invalid integer: expected %d, got %d", 0, v.Integer())
	}
	v, rn, err = r.ReadValue()
	sum += rn
	if err != io.EOF {
		t.Fatalf("invalid error: expected %v, got %v", io.EOF, err)
	}
	if sum != len(data) {
		t.Fatalf("invalid read count: expected %d, got %d", len(data), sum)
	}
}

func TestFloats(t *testing.T) {
	sum := 0
	data := []byte(":1234567\r\n+-90898\r\n$6\r\n12.345\r\n-90284.987\r\n")
	r := NewRESPReader(bytes.NewBuffer(data))
	v, rn, err := r.ReadValue()
	sum += rn
	if err != nil {
		t.Fatal(err)
	}
	if v.Float() != 1234567 {
		t.Fatalf("invalid integer: expected %v, got %v", 1234567, v.Float())
	}
	v, rn, err = r.ReadValue()
	sum += rn
	if err != nil {
		t.Fatal(err)
	}
	if v.Float() != -90898 {
		t.Fatalf("invalid integer: expected %v, got %v", -90898, v.Float())
	}
	v, rn, err = r.ReadValue()
	sum += rn
	if err != nil {
		t.Fatal(err)
	}
	if v.Float() != 12.345 {
		t.Fatalf("invalid integer: expected %v, got %v", 12.345, v.Float())
	}
	v, rn, err = r.ReadValue()
	sum += rn
	if err != nil {
		t.Fatal(err)
	}
	if v.Float() != 90284.987 {
		t.Fatalf("invalid integer: expected %v, got %v", 90284.987, v.Float())
	}
	v, rn, err = r.ReadValue()
	sum += rn
	if err != io.EOF {
		t.Fatalf("invalid error: expected %v, got %v", io.EOF, err)
	}
	if sum != len(data) {
		t.Fatalf("invalid read count: expected %d, got %d", len(data), sum)
	}
}

// TestLotsaRandomness does generates N resp messages and reads the values though a Reader.
// It then marshals the values back to strings and compares to the original.
// All data and resp types are random.
func TestLotsaRandomness(t *testing.T) {
	n := 10000
	var anys []string
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		any := randRESPAny()
		anys = append(anys, any)
		buf.WriteString(any)
	}
	r := NewRESPReader(bytes.NewBuffer(buf.Bytes()))
	for i := 0; i < n; i++ {
		v, _, err := r.ReadValue()
		if err != nil {
			t.Fatal(err)
		}
		ts := fmt.Sprintf("%v", v.Type())
		if ts == "Unknown" {
			t.Fatal("got 'Unknown'")
		}
		tvs := fmt.Sprintf("%v %v %v %v %v %v %v %v",
			v.String(), v.Float(), v.Integer(), v.Array(),
			v.Bool(), v.Bytes(), v.IsNull(), v.Error(),
		)
		if len(tvs) < 10 {
			t.Fatal("conversion error")
		}
		if !v.Equals(v) {
			t.Fatal("equals failed")
		}
		resp, err := v.MarshalRESP()
		if err != nil {
			t.Fatal(err)
		}
		if string(resp) != anys[i] {
			t.Fatalf("resp failed to remarshal #%d\n-- original --\n%s\n-- remarshalled --\n%s\n-- done --", i, anys[i], string(resp))
		}
	}
}

func TestBigFragmented(t *testing.T) {
	b := make([]byte, 10*1024*1024)
	if _, err := rand.Read(b); err != nil {
		t.Fatal(err)
	}
	cmd := []byte("*3\r\n$3\r\nSET\r\n$3\r\nKEY\r\n$" + strconv.FormatInt(int64(len(b)), 10) + "\r\n" + string(b) + "\r\n")
	cmdlen := len(cmd)
	pr, pw := io.Pipe()
	frag := 1024
	go func() {
		defer pw.Close()
		for len(cmd) >= frag {
			if _, err := pw.Write(cmd[:frag]); err != nil {
				t.Error(err)
				return
			}
			cmd = cmd[frag:]
		}
		if len(cmd) > 0 {
			if _, err := pw.Write(cmd); err != nil {
				t.Error(err)
				return
			}
		}
	}()
	r := NewRESPReader(pr)
	value, n, err := r.readArrayValue()
	if err != nil {
		t.Fatal(err)
	}
	if n != cmdlen {
		t.Fatalf("expected %v, got %v", cmdlen, n)
	}
	arr := value.Array()
	if len(arr) != 3 {
		t.Fatalf("expected 3, got %v", len(arr))
	}
	if arr[0].String() != "SET" {
		t.Fatalf("expected 'SET', got %v", arr[0].String())
	}
	if arr[1].String() != "KEY" {
		t.Fatalf("expected 'KEY', got %v", arr[0].String())
	}
	if bytes.Compare(arr[2].Bytes(), b) != 0 {
		t.Fatal("bytes not equal")
	}
}
