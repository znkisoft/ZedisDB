package main

import (
	"bytes"
	"log"
	"strconv"
)

type Message []byte

type Decoder interface {
	DecodeSimpleString(message Message) (string, error)
	DecodeBulkString(message Message) (string, error)
	DecodeArray(message Message) ([]string, error)
	CheckType(message Message) (uint8, error)
}
type RESPDecode struct{}

type SimpleString string
type BulkString string
type Array []string

// zedis errors
const (
	ZedisErrIo = iota + 1
	ZedisErrOther
	ZedisErrEof
	ZedisErrProtocol
	ZedisErrOom
	ZedisErrTimeout
)

// zedis reply type
const (
	ZedisReplyString = iota + 1
	ZedisReplyArray
	ZedisReplyInteger
	ZedisReplyNil
	ZedisReplyStatus
	ZedisReplyError
	ZedisReplyDouble
	ZedisReplyBool
	ZedisReplyMap
	ZedisReplySet
	ZedisReplyAttr
	ZedisReplyPush
	ZedisReplyBignum
	ZedisReplyVerb
	ZedisReplyBulkString // Bulk Strings are used in order to represent a single binary safe string up to 512 MB in length.
	ZedisReplyUnknown
)

var (
	MsgConnected = []byte("*1\r\n$7\r\nCOMMAND\r\n")
	MsgPing      = []byte("*1\r\n$4\r\nCOMMAND\r\n")
	MsgPong      = []byte("+PONG\r\n")

	delimiter = []byte("\r\n")
)

func NewMessage(m []byte, length int) Message {
	// TODO remove
	log.Printf("[data](length:%d): %v\n************************", length, string(m))
	return m
}

/*
CheckType
	For Simple Strings the first byte of the reply is "+"
	For Errors the first byte of the reply is "-"
	For Integers the first byte of the reply is ":"
	For Bulk Strings the first byte of the reply is "$"
	For Arrays the first byte of the reply is "*"
*/
func (d RESPDecode) CheckType(m Message) uint8 {
	if bytes.HasPrefix(m, []byte{'+'}) {
		return ZedisReplyString
	} else if bytes.HasPrefix(m, []byte{'-'}) {
		return ZedisReplyError
	} else if bytes.HasPrefix(m, []byte{':'}) {
		return ZedisReplyInteger
	} else if bytes.HasPrefix(m, []byte{'$'}) {
		return ZedisReplyBulkString
	} else if bytes.HasPrefix(m, []byte{'*'}) {
		return ZedisReplyArray
	} else {
		return ZedisReplyUnknown
	}
}

func (d *RESPDecode) DecodeSimpleString(msg Message) (string, error) {
	length := len(msg)

	if length < 3 {
		return "", Err{code: InCompleteMessage}
	}

	if bytes.Equal(msg[length-2:length], delimiter) {
		return string(msg[1 : length-2]), nil
	}
	return "", Err{code: NotImplement}
}

func (d *RESPDecode) DecodeBulkString(msg Message) (string, error) {
	length := len(msg)

	// minimum length: "$-1\r\n"
	if length < 4 {
		return "", Err{code: InCompleteMessage}
	}

	if bytes.Equal(msg[length-2:length], delimiter) {
		// byte num to int with explicit cast
		count := int(msg[1] - '0')

		// "$-1\r\n"
		if count == '-' {
			// TODO implement Null Bulk String
			return "", Err{code: InCompleteMessage}
		}

		// "$0\r\n\r\n"
		if count == 0 {
			return "", nil
		}

		array := bytes.Split(msg, delimiter)
		if len(array) == 3 {
			return string(array[1]), nil
		}

		return "", Err{code: InCompleteMessage}
	}
	return "", Err{code: InCompleteMessage}
}

func (d *RESPDecode) DecodeArray(msg Message) ([]string, error) {
	length := len(msg)
	var strArray []string

	if length < 9 {
		return nil, Err{code: InCompleteMessage}
	}

	if bytes.Equal(msg[length-2:], delimiter) {
		array := bytes.Split(msg[0:length-2], delimiter)
		if len(array) < 3 {
			return nil, Err{code: InCompleteMessage}
		}

		count, _ := strconv.Atoi(string(array[0][1:]))
		filtered := func(s []byte) bool { return !bytes.HasPrefix(s, []byte("$")) }
		for _, byteArray := range array[1:] {
			if filtered(byteArray) {
				strArray = append(strArray, string(byteArray))
			}
		}

		if count != len(strArray) {
			return nil, Err{code: InCompleteMessage}
		}
		return strArray, nil

	}
	return nil, Err{code: InCompleteMessage}
}
