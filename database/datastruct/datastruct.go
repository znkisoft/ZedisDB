package datastruct

type Dictionary interface {
	Set(key string, value *ZedisObject) error
	Delete(key string)
	Get(key string) (*ZedisObject, bool)
	Keys() []string
	Size() int
	Clear()
}
