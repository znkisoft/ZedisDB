package database

import (
	"github.com/znkisoft/zedisDB/database/datastruct"
)

type Db struct {
	Dict    *datastruct.Dict
	Expires *datastruct.Dict
}

func NewDb() *Db {
	debugDict := datastruct.NewDict()
	_ = debugDict.Set("version", "debug")
	_ = debugDict.Set("num", 1)
	return &Db{
		Dict:    debugDict,
		Expires: datastruct.NewDict(),
	}
}
