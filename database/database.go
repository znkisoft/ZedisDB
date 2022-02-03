package database

import (
	"github.com/znkisoft/zedisDB/database/datastruct"
)

type Db struct {
	Dict    *datastruct.Dict
	Expires *datastruct.Dict
}

func NewDb() *Db {
	return &Db{
		Dict:    datastruct.NewDict(),
		Expires: datastruct.NewDict(),
	}
}
