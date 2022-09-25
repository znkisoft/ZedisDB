package datastruct

import (
	"github.com/znkisoft/zedisDB/lru"
)

type ObjectType uint8

const (
	StringTyp ObjectType = iota
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
	typ         ObjectType
	encodingTyp EncodingType
	lru         int64 // unix millisecond
	ptr         interface{}
}

func CreateZedisObject(t ObjectType, ptr interface{}) *ZedisObject {
	return &ZedisObject{
		typ:         t,
		encodingTyp: RawEncodingTyp,
		lru:         lru.GetLRUClock(),
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

func (o *ZedisObject) Typ() ObjectType {
	return o.typ
}

func (o *ZedisObject) ObjectIdleTime() int64 {
	return lru.GetLRUClock() - o.lru
}
