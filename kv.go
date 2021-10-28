package main

import "sync"

type KV struct {
	mu   sync.Mutex
	data map[string]interface{}
}

func New() *KV {
	return &KV{
		data: make(map[string]interface{}),
	}
}

func (kv *KV) Get(key string) (interface{}, bool) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	if v, ok := kv.data[key]; ok {
		return v, true
	}
	return nil, false
}

func (kv *KV) Set(key string, value interface{}) {
	kv.mu.Lock()
	kv.data[key] = value
	defer kv.mu.Unlock()
}
