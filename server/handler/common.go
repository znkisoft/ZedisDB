package handler

import (
	"github.com/znkisoft/zedisDB/parser"
)

func resolveCmdArgs(args []parser.Value, size int) (string, []parser.Value, error) {
	if len(args) < size {
		return "", nil, parser.Err{Type: parser.Client, Message: "wrong number of arguments"}
	}
	return args[0].String(), args[1:], nil
}
