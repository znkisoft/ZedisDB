package main

import (
	"bytes"
	"errors"
	"log"
)

type Message []byte

type Decoder interface {
	DecodeSimpleString(message Message) (string, error)
	DecodeBulkString(message Message) ([]string, error)
	DecodeArray(message Message) (string, error)
	CheckType(message Message) (uint8, error)
}
type RESPDecode struct{}

// redis errors
const (
	REDIS_ERR_IO = iota + 1
	REDIS_ERR_OTHER
	REDIS_ERR_EOF
	REDIS_ERR_PROTOCOL
	REDIS_ERR_OOM
	REDIS_ERR_TIMEOUT
)

// redis reply type
const (
	RedisReplyString = iota + 1
	RedisReplyArray
	RedisReplyInteger
	RedisReplyNil
	RedisReplyStatus
	RedisReplyError
	RedisReplyDouble
	RedisReplyBool
	RedisReplyMap
	RedisReplySet
	RedisReplyAttr
	RedisReplyPush
	RedisReplyBignum
	RedisReplyVerb
	RedisReplyBulkString // Bulk Strings are used in order to represent a single binary safe string up to 512 MB in length.
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
func (d RESPDecode) CheckType(m Message) (uint8, error) {
	if bytes.HasPrefix(m, []byte{'+'}) {
		return RedisReplyString, nil
	} else if bytes.HasPrefix(m, []byte{'-'}) {
		return RedisReplyError, nil
	} else if bytes.HasPrefix(m, []byte{':'}) {
		return RedisReplyInteger, nil
	} else if bytes.HasPrefix(m, []byte{'$'}) {
		return RedisReplyBulkString, nil
	} else if bytes.HasPrefix(m, []byte{'*'}) {
		return RedisReplyArray, nil
	}
	return 0, errors.New("message's prefix is not supported")
}

func (d *RESPDecode) DecodeSimpleString(msg Message) (string, error) {
	length := len(msg)

	if length < 3 {
		return "", Err{code: InCompleteStr}
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
		return "", Err{code: InCompleteStr}
	}

	if bytes.Equal(msg[length-2:length], delimiter) {
		// byte num to int with explicit cast
		count := int(msg[1] - '0')

		// "$-1\r\n"
		if count == '-' {
			// TODO implement Null Bulk String
			return "", Err{code: InCompleteStr}
		}

		// "$0\r\n\r\n"
		if count == 0 {
			return "", nil
		}

		byteCollection := bytes.Split(msg, delimiter)
		if len(byteCollection) == 3 {
			return string(byteCollection[1]), nil
		}

		return "", Err{code: InCompleteStr}
	}
	return "", Err{code: InCompleteStr}
}

func (d *RESPDecode) DecodeArray(msg Message) ([]string, error) {
	return []string{}, Err{code: NotImplement}
}
