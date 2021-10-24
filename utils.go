package main

import (
	"fmt"
)

func CheckError(err error) {
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s\n", err))
	}
}
