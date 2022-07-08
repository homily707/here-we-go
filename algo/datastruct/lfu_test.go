package datastruct

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	var cache Cache = NewLFUCache()
	cache.Put(1, "a")
	cache.Put(2, "b")
	cache.Put(3, "c")
	cache.Put(4, "d")
	cache.Put(5, "e")
	fmt.Println(cache.String())
	cache.Put(6, "f")
	fmt.Println(cache.String())
	cache.Put(6, "ff")
	cache.Put(6, "fff")
	cache.Put(1, "a")
	cache.Put(2, "b")
	fmt.Println(cache.String())
	cache.Put(4, "d")
	cache.Put(5, "e")
	fmt.Println(cache.String())
}
