package parser

import (
	"bytes"
	"testing"
)

func TestWriter(t *testing.T) {
	var buf bytes.Buffer
	wr := NewRESPWriter(&buf)
	wr.WriteArray(MultiBulkValue("HELLO", 1, 2, 3).Array())
	wr.WriteBytes([]byte("HELLO"))
	wr.WriteString("HELLO")
	wr.WriteSimpleString("HELLO")
	wr.WriteError(ErrProtocol{
		Type:    Server,
		Message: "HELLO",
	})
	wr.WriteInteger(1)
	wr.WriteNull()
	wr.WriteValue(SimpleStringValue("HELLO"))

	res := "" +
		"*4\r\n$5\r\nHELLO\r\n$1\r\n1\r\n$1\r\n2\r\n$1\r\n3\r\n" +
		"$5\r\nHELLO\r\n" +
		"$5\r\nHELLO\r\n" +
		"+HELLO\r\n" +
		"-server HELLO\r\n" +
		":1\r\n" +
		"$-1\r\n" +
		"+HELLO\r\n"
	if buf.String() != res {
		t.Fatalf("expected '%v', got '%v'", res, buf.String())
	}

}
