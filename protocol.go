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
	DecodeArray(message Message) ([]string, error)
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
)

var (
	MsgConnected = []byte("*1\r\n$7\r\nCOMMAND\r\n")
	MsgPing      = []byte("*1\r\n$4\r\nCOMMAND\r\n")
	MsgPong      = []byte("+PONG\r\n")
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
		return RedisReplyArray, nil
	} else if bytes.HasPrefix(m, []byte{'*'}) {
		return RedisReplyArray, nil
	}
	return 0, errors.New("message's prefix is not supported")
}

func (d *RESPDecode) DecodeSimpleString(msg Message) (string, error) {
	return "", Err{code: InCompleteStr}
}

func (d *RESPDecode) DecodeBulkString(msg Message) ([]string, error) {
	return []string{}, Err{code: NotImplement}
}

func (d *RESPDecode) DecodeArray(msg Message) ([]string, error) {
	return []string{}, Err{code: NotImplement}
}
