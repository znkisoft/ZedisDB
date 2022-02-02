package datastruct

import (
	"testing"
)

func TestDict(t *testing.T) {
	dict := NewDict()

	t.Run("TestGetNull", func(t *testing.T) {
		if v, _ := dict.Get("test"); v != nil {
			t.Error("Expected nil")
		}
	})

	t.Run("TestSetAndGet", func(t *testing.T) {
		err := dict.Set("key", "value")
		if err != nil {
			t.Log(err)
		}
		if v, _ := dict.Get("key"); v != "value" {
			t.Errorf("Expected value: \"value\"")
		}
	})

	t.Run("TestDelete", func(t *testing.T) {
		err := dict.Set("key2", "value")
		if err != nil {
			t.Log(err)
		}
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
	// goos: darwin
	// goarch: amd64
	// pkg: github.com/znkisoft/zedisDB/datastruct
	// cpu: VirtualApple @ 2.50GHz
	// BenchmarkKV
	// BenchmarkKV-8   	 7026129	       152.9 ns/op
	for i := 0; i < b.N; i++ {
		dict := NewDict()
		_ = dict.Set("key", "value")
		dict.Get("key")
	}
}
