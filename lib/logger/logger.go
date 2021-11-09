package logger

import (
	"log"
	"os"
)

var (
	CommonLog *log.Logger
	ErrorLog  *log.Logger
)

func init() {
	CommonLog = log.New(os.Stdout, "[notification]:", log.LstdFlags)
	ErrorLog = log.New(os.Stderr, "[error]:", log.LstdFlags)
}
