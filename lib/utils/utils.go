package utils

import (
	"errors"
	"fmt"
	"io"
)

func CheckError(err error) {
	if err != nil {
		if errors.Is(err, io.EOF) {
			return
		}
		panic(fmt.Sprintf("[Error]: %s\n", err))
	}
}
