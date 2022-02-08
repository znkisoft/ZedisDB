package datastruct

import (
	"reflect"
	"testing"
)

func TestDict(t *testing.T) {
	dict := NewDict()
	o := CreateZedisObject(StringTyp, "value")

	t.Run("TestGetNull", func(t *testing.T) {
		if v, _ := dict.Get("test"); v != nil {
			t.Error("Expected nil")
		}
	})

	t.Run("TestSetAndGet", func(t *testing.T) {
		err := dict.Set("key", o)
		if err != nil {
			t.Log(err)
		}
		if v, _ := dict.Get("key"); !reflect.DeepEqual(v, o) {
			t.Errorf("Expected value: \"value\"")
		}
	})

	t.Run("TestDelete", func(t *testing.T) {
		err := dict.Set("key2", o)
		if err != nil {
			t.Log(err)
		}
		if v, _ := dict.Get("key2"); !reflect.DeepEqual(v, o) {
			t.Errorf("Expected value: \"value\"")
		}
		dict.Delete("key")
		if v, _ := dict.Get("key"); v != nil {
			t.Errorf("Expected value: nil")
		}
	})
}

func BenchmarkKV(b *testing.B) {
	o := CreateZedisObject(StringTyp, "value")
	// goos: darwin
	// goarch: arm64
	// pkg: github.com/znkisoft/zedisDB/database/datastruct
	// BenchmarkKV
	// BenchmarkKV-8   	11695563	   	   101.2 ns/op
	for i := 0; i < b.N; i++ {
		dict := NewDict()
		_ = dict.Set("key", o)
		dict.Get("key")
	}
}
