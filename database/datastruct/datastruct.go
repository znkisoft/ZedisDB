package datastruct

import (
	"github.com/znkisoft/zedisDB/database/container"
)

type Dictionary interface {
	Set(key string, value *container.ZedisObject) error
	Delete(key string)
	Get(key string) (*container.ZedisObject, bool)
	Keys() []string
	Size() int
	Clear()
}
