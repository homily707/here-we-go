package datastruct

import "container/list"

type Cache [T any] interface {
	Add(key, value T) bool

	Get(Key T) (value T, ok bool)
}

type LRUCache struct {
	data map[any] any
	freq map[int] *list.List
	elements map[any]*list.Element
	lf   int
	cap  int
	len  int
}

type Node struct {
	key   any
	value any
	freq  int
}

func NewLRUCache(cap int) *LRUCache{
	return &LRUCache{
		data: map[any]any{},
		freq: map[int]*list.List,
		elements: map[any]*list.Element,
		lf:   0,
		cap:  cap,
		len:  0,
	}
}

func (L *LRUCache) Add(key, value any) bool {
	var e *list.Element

	if _,ok := L.data[key]; ok != true {
		L.data[key] = value
		L.len ++
		e = L.freq[1].PushFront(key)
		L.elements[key] = e
	} else {
		L.data[key] = value
		e = L.elements[key]

	}

	return true
}

func (L *LRUCache) Get(Key any) (value any, ok bool) {

}

//type LRUCache [K any, T any] struct {
//	data map[K]T
//	freq map[int]list.List
//	lf   int
//	cap  int
//}
//
//func NewLRUCache [K any, T any] (cap int) {
//	return LRUCache [K,T] {
//		data: map[K]T,
//		freq: nil,
//		lf:   0,
//		cap:  0,
//	}
//}
//
//func (L LRUCache[K,T]) Add(key K, value T) bool {
//
//}
//
//func (L LRUCache[T]) Get(Key T) (value T, ok bool) {
//
//}

