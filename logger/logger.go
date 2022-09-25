package logger

import (
	"log"
	"os"
)

var (
	CommonLog = log.New(os.Stdout, "[notification]:", log.LstdFlags)
	ErrorLog  = log.New(os.Stderr, "[error]:", log.LstdFlags)
)
