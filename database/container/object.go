package container

import (
	"github.com/znkisoft/zedisDB/pkg"
)

type Type uint8

const (
	StringTyp Type = iota
	HashTyp
	ListTyp
	SetTyp
	ZSetTyp
)

type EncodingType uint8

const (
	RawEncodingTyp EncodingType = iota
)

type ZedisObject struct {
	typ         Type
	encodingTyp EncodingType
	lru         int64 // unix millisecond
	ptr         interface{}
}

func CreateZedisObject(t Type, ptr interface{}) *ZedisObject {
	return &ZedisObject{
		typ:         t,
		encodingTyp: RawEncodingTyp,
		lru:         pkg.GetLRUClock(),
		ptr:         ptr,
	}
}

func (o *ZedisObject) EncodingTyp() EncodingType {
	return o.encodingTyp
}

func (o *ZedisObject) Lru() int64 {
	return o.lru
}

func (o *ZedisObject) Ptr() interface{} {
	return o.ptr
}

func (o *ZedisObject) Typ() Type {
	return o.typ
}

func (o *ZedisObject) ObjectIdleTime() int64 {
	return pkg.GetLRUClock() - o.lru
}
