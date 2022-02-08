package datastruct

import (
	"sync"
)

type BaseDict struct {
	mu   sync.RWMutex
	data map[string]*ZedisObject
}

func NewDict() *BaseDict {
	return &BaseDict{
		data: make(map[string]*ZedisObject),
	}
}

func (dict *BaseDict) Get(key string) (*ZedisObject, bool) {
	dict.mu.RLock()
	defer dict.mu.RUnlock()
	if v, ok := dict.data[key]; ok {
		return v, true
	}
	return nil, false
}

func (dict *BaseDict) Set(key string, value *ZedisObject) error {
	dict.mu.Lock()
	dict.data[key] = value
	defer dict.mu.Unlock()
	return nil
}

func (dict *BaseDict) Delete(key string) {
	dict.mu.Lock()
	delete(dict.data, key)
	defer dict.mu.Unlock()
}

func (dict *BaseDict) Keys() []string {
	dict.mu.RLock()
	defer dict.mu.RUnlock()
	if len(dict.data) == 0 {
		return []string{}
	}
	keys := make([]string, 0, len(dict.data))
	for k := range dict.data {
		keys = append(keys, k)
	}
	return keys
}

func (dict *BaseDict) Size() int {
	return len(dict.data)
}

func (dict *BaseDict) Clear() {
	dict.mu.Lock()
	dict.data = make(map[string]*ZedisObject)
	dict.mu.Unlock()
}
