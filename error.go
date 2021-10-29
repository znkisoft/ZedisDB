package main

import (
	"errors"
	"fmt"
)

type Err struct {
	msg  string
	code int
}

const (
	NotImplement      = -1
	InCompleteMessage = iota
)

var (
	errMessageMap = map[int]string{
		NotImplement:      "not implement yet",
		InCompleteMessage: "incomplete message",
	}
	ErrUnknownCommand = errors.New("err command unknown")
)

func (e Err) Error() string {
	return fmt.Sprintf("error[%d]: %s", e.code, getMsg(e.code))
}

func getMsg(code int) string {
	msg, ok := errMessageMap[code]
	if ok {
		return msg
	}
	return errMessageMap[NotImplement]
}
