package database

import (
	"github.com/znkisoft/zedisDB/database/container"
	"github.com/znkisoft/zedisDB/database/datastruct"
)

type Db struct {
	// ID      int
	dict    *datastruct.BaseDict
	expires *datastruct.BaseDict
}

func NewDb() *Db {
	return &Db{
		dict:    datastruct.NewDict(),
		expires: datastruct.NewDict(),
	}
}

func (db *Db) Set(key string, value *container.ZedisObject) error {
	return db.dict.Set(key, value)
}

func (db *Db) Get(key string) (*container.ZedisObject, bool) {
	return db.dict.Get(key)
}

func (db *Db) Del(key string) error {
	db.dict.Delete(key)
	return nil
}

// func (db *Db) GetExpire(key string) {
//
// }
//
// func (db *Db) SetExpire(key string) {
//
// }
//
// func (db *Db) RemoveExpire(key string) {
//
// }
//
// func (db *Db) PropagateExpire(key string) {
//
// }
//
// func (db *Db) ExpireIfNeeded(key string) {
//
// }
//
// func (db *Db) dbAdd(key string) {
//
// }
//
// func (db *Db) dbOverwrite(key string) {
//
// }
//
// func (db *Db) dbExists(key string) {
//
// }
//
// func (db *Db) dbRandomKey(key string) {
//
// }
//
// func (db *Db) EmptyDb(key string) {
//
// }
