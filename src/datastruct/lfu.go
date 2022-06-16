package datastruct

import "container/list"

type LRUCache interface {
	Add(key, value interface{}) bool

	Get(Key interface{}) (value interface{}, ok bool)
}

type LRU struct {
	list.List
}
