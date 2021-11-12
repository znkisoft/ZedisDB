package datastructure

import (
	"sync"
)

type Dict struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

func NewDict() *Dict {
	return &Dict{
		data: make(map[string]interface{}),
	}
}

func (dict *Dict) Get(key string) (interface{}, bool) {
	dict.mu.Lock()
	defer dict.mu.Unlock()
	if v, ok := dict.data[key]; ok {
		return v, true
	}
	return nil, false
}

func (dict *Dict) Set(key string, value interface{}) {
	dict.mu.Lock()
	dict.data[key] = value
	defer dict.mu.Unlock()
}

func (dict *Dict) Delete(key string) {
	dict.mu.Lock()
	delete(dict.data, key)
	defer dict.mu.Unlock()
}
