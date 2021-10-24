package main

import (
	"fmt"
)

type Err struct {
	msg  string
	code int
}

const (
	NotImplement  = -1
	InCompleteStr = iota + 1
)

var errMessageMap = map[int]string{
	NotImplement:  "Not Implement yet",
	InCompleteStr: "incomplete string",
}

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
