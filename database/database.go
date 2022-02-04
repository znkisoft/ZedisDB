package database

import (
	"github.com/znkisoft/zedisDB/database/container"
	"github.com/znkisoft/zedisDB/database/datastruct"
)

type ZedisObject struct {
	Type container.Type
	Ptr  interface{}
}

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

func (db *Db) SetKey(key string, value ZedisObject) {

}

func (db *Db) LookupKey(key string) *ZedisObject {
	return nil
}

// TODO LRU eviction policy
// eviction rules
// - [x] expired keys expire after lookup
// - [ ] shortage of memory
// - [ ] cron job to run eviction
