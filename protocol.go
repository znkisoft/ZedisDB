package main

import (
	"bytes"
	"errors"
	"log"
)

type Messager interface {
	String() string
}

type Message []byte

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

func (m Message) String() string {
	return string(m)
}

func (m Message) BulkStringParse() [][]byte {
	if len(m) == 0 {
		return [][]byte{}
	}
	return nil
}

/*
CheckType
	For Simple Strings the first byte of the reply is "+"
	For Errors the first byte of the reply is "-"
	For Integers the first byte of the reply is ":"
	For Bulk Strings the first byte of the reply is "$"
	For Arrays the first byte of the reply is "*"
*/
func (m Message) CheckType() (uint8, error) {
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
