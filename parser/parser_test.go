package parser

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"testing"
)

func TestAnyValues(t *testing.T) {
	var vs = []interface{}{
		nil,
		10, uint(10), int8(10),
		uint8(10), int16(10), uint16(10),
		int32(10), uint32(10), int64(10),
		uint64(10), bool(true), bool(false),
		float32(10), float64(10),
		[]byte("hello"), string("hello"),
	}
	for i, v := range vs {
		if AnyValue(v).String() == "" && v != nil {
			t.Fatalf("missing string value for #%d: '%v'", i, v)
		}
	}
}

func TestMarshalStrangeValue(t *testing.T) {
	var v Value
	v.null = true
	b, err := marshalAnyRESP(v)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "$-1\r\n" {
		t.Fatalf("expected '%v', got '%v'", "$-1\r\n", string(b))
	}
	v.null = false

	_, err = marshalAnyRESP(v)
	if err == nil || err.Error() != "unknown resp type" {
		t.Fatalf("expected '%v', got '%v'", "unknown resp type encountered", err)
	}
}

func randRESPInteger() string {
	return fmt.Sprintf(":%d\r\n", (randInt()%1000000)-500000)
}
func randRESPSimpleString() string {
	return "+" + strings.Replace(randString(), "\r\n", "", -1) + "\r\n"
}
func randRESPError() string {
	return "-" + strings.Replace(randString(), "\r\n", "", -1) + "\r\n"
}
func randRESPBulkString() string {
	s := randString()
	if len(s)%1024 == 0 {
		return "$-1\r\n"
	}
	return "$" + strconv.FormatInt(int64(len(s)), 10) + "\r\n" + s + "\r\n"
}
func randRESPArray() string {
	n := randInt() % 10
	if n%10 == 0 {
		return "$-1\r\n"
	}
	s := "*" + strconv.FormatInt(int64(n), 10) + "\r\n"
	for i := 0; i < n; i++ {
		rn := randInt() % 100
		if rn == 0 {
			s += randRESPArray()
		} else {
			switch (rn - 1) % 4 {
			case 0:
				s += randRESPInteger()
			case 1:
				s += randRESPSimpleString()
			case 2:
				s += randRESPError()
			case 3:
				s += randRESPBulkString()
			}
		}
	}
	return s
}

func randInt() int {
	n := int(binary.LittleEndian.Uint64(randBytes(8)))
	if n < 0 {
		n *= -1
	}
	return n
}

func randBytes(n int) []byte {
	b := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		log.Fatal("random error: " + err.Error())
	}
	return b
}

func randString() string {
	return string(randBytes(randInt() % 1024))
}

func randRESPAny() string {
	switch randInt() % 5 {
	case 0:
		return randRESPInteger()
	case 1:
		return randRESPSimpleString()
	case 2:
		return randRESPError()
	case 3:
		return randRESPBulkString()
	case 4:
		return randRESPArray()
	}
	panic("?")
}

func BenchmarkRead(b *testing.B) {
	n := 1000
	var buf bytes.Buffer
	for k := 0; k < n; k++ {
		buf.WriteString(randRESPAny())
	}
	bb := buf.Bytes()
	b.ResetTimer()
	var j int
	var r *RESPReader
	// start := time.Now()
	var k int
	for i := 0; i < b.N; i++ {
		if j == 0 {
			r = NewRESPReader(bytes.NewBuffer(bb))
			j = n
		}
		_, _, err := r.ReadValue()
		if err != nil {
			b.Fatal(err)
		}
		j--
		k++
	}
	// fmt.Printf("\n%f\n", float64(k)/(float64(time.Now().Sub(start))/float64(time.Second)))
	// goos: darwin
	// goarch: amd64
	// pkg: github.com/znkisoft/zedisDB/parser
	// cpu: VirtualApple @ 2.50GHz
	// BenchmarkRead
	// BenchmarkRead-8   	 1588827	       724.7 ns/op
}
