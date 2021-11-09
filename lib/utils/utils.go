package utils

import (
	"errors"
	"io"

	"github.com/znkisoft/zedisDB/lib/logger"
)

func CheckError(err error) {
	if err != nil {
		if errors.Is(err, io.EOF) {
			return
		}
		logger.ErrorLog.Fatal(err)
	}
}
