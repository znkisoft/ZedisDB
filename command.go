package main

import (
	"fmt"
	"strings"
)

type Commnader interface {
	ping() []byte
	echo() []byte
	HandleCommand()
}

type Command struct {
	Action string
	Args   []string
}

func (c *Command) ping() []byte {
	return []byte("+pong\r\n")
}

func (c *Command) echo(arg string) []byte {
	return []byte(fmt.Sprintf("+%s\r\n", arg))
}

func (c *Command) HandleCommand(action string, args []string) []byte {
	if strings.ToLower(action) == "ping" {
		return c.ping()
	}

	if strings.ToLower(action) == "echo" && len(args) == 1 {
		return c.echo(args[0])
	}

	return []byte{0}
}

func (c *Command) errorCommand(redisErrType int) []byte {
	return []byte{0}
}
