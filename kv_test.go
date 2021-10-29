package main

import (
	"testing"
)

func TestKV(t *testing.T) {
	// test kv functionality
	kv := New()

	t.Run("TestGetNull", func(t *testing.T) {
		if v, _ := kv.Get("test"); v != nil {
			t.Error("Expected nil")
		}
	})

	t.Run("TestSetAndGet", func(t *testing.T) {
		kv.Set("key", "value")
		if v, _ := kv.Get("key"); v != "value" {
			t.Errorf("Expected value: \"value\"")
		}
	})

	t.Run("TestDelete", func(t *testing.T) {
		kv.Set("key2", "value")
		if v, _ := kv.Get("key2"); v != "value" {
			t.Errorf("Expected value: \"value\"")
		}
		kv.Delete("key")
		if v, _ := kv.Get("key"); v != nil {
			t.Errorf("Expected value: nil")
		}
	})
}

func BenchmarkKV(b *testing.B) {

	for i := 0; i < b.N; i++ {
		kv := New()
		kv.Set("key", "value")
		kv.Get("key")
	}
}
