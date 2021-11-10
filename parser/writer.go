package parser

import (
	"io"
)

type RESPWriter struct {
	R io.Writer
}

func NewRESPWriter(w io.Writer) *RESPWriter {
	return &RESPWriter{
		R: w,
	}
}
