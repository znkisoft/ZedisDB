package main

import (
	"testing"
)

func TestKV(t *testing.T) {
	// test kv functionality
	kv := New()

	t.Run("TestGetNull", func(t *testing.T) {
		if v,_:=kv.Get("test") ; v!= nil {
			t.Error("Expected nil")
		}
	})

	t.Run("TestSetAndGet", func(t *testing.T) {
		kv.Set("key", "value")
		if v, _ := kv.Get("key"); v != "value" {
			t.Errorf("Expected value: \"value\"")
		}
	})


	
}
