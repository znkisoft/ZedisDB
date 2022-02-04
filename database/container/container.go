package container

type Container interface{}

type Type uint8

const (
	HashTyp Type = iota
	ListTyp
	SetTyp
	ZSetTyp
	StringTyp
)
