package lru

import "testing"

func TestCache_Add(t *testing.T) {
	lru := New(int64(10), nil)
	key1 := "s"
	val1 := String("b")
	lru.Add(key1, val1)

	if v, ok := lru.Get(key1); ok {
		if v.Len() != val1.Len() {
			t.Fatalf("cache value not equal source value")
		}
	}
}
