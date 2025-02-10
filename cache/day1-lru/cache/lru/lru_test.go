package lru

import (
	"log"
	"testing"
	"unsafe"
)

type String string

func (s String) Len() int {
	return len(s)
}

type StringSlice []string

func (s StringSlice) Len() int {
	size := int(unsafe.Sizeof(s)) // Size of the slice header

	for _, str := range s {
		size += len(str) + int(unsafe.Sizeof(str))
	}

	return size
}

func TestCache_Add(t *testing.T) {
	c := New(2000, func(key string, v Value) {
		log.Printf("del success %v => %v ", key, v)
	})
	k1 := "k1"
	test1 := StringSlice([]string{"WOO", "JACK"})
	log.Printf("test1 length %v", test1.Len())
	c.Add(k1, test1)
	log.Printf("缓存内存 %v", c.Len())

	if v, ok := c.Get(k1); ok {
		log.Printf("get %v => %v", k1, v)
		if v.Len() != test1.Len() {
			t.Fatal("缓存的值不对")
		}
	} else {
		t.Fatal("get fail")
	}

}
