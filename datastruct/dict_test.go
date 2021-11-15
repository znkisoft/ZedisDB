package datastruct

import (
	"testing"
)

func TestDict(t *testing.T) {
	// test dict functionality
	dict := NewDict()

	t.Run("TestGetNull", func(t *testing.T) {
		if v, _ := dict.Get("test"); v != nil {
			t.Error("Expected nil")
		}
	})

	t.Run("TestSetAndGet", func(t *testing.T) {
		dict.Set("key", "value")
		if v, _ := dict.Get("key"); v != "value" {
			t.Errorf("Expected value: \"value\"")
		}
	})

	t.Run("TestDelete", func(t *testing.T) {
		dict.Set("key2", "value")
		if v, _ := dict.Get("key2"); v != "value" {
			t.Errorf("Expected value: \"value\"")
		}
		dict.Delete("key")
		if v, _ := dict.Get("key"); v != nil {
			t.Errorf("Expected value: nil")
		}
	})
}

func BenchmarkKV(b *testing.B) {

	for i := 0; i < b.N; i++ {
		dict := NewDict()
		dict.Set("key", "value")
		dict.Get("key")
	}
}
