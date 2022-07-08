package datastruct

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type Cache interface {
	fmt.Stringer
	Put(key int, value interface{})
	Get(key int) interface{}
	Len() int
}

type LFUCache struct {
	data     Cache
	minFreq  int
	capacity int
	freqList map[int]*list.List
	elements map[int]*list.Element
}

type node struct {
	key   int
	value interface{}
	freq  int
}

func NewLFUCache() LFUCache {
	return LFUCache{
		data:     nil,
		minFreq:  0,
		capacity: 5,
		freqList: map[int]*list.List{},
		elements: map[int]*list.Element{},
	}
}

func (lfu LFUCache) String() string {
	builder := strings.Builder{}
	for k, v := range lfu.freqList {
		builder.WriteString(strconv.Itoa(k) + ": ")
		for e := v.Front(); e != nil; e = e.Next() {
			builder.WriteString(fmt.Sprintf("%v", e.Value.(node)))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (lfu LFUCache) Get(key int) interface{} {
	e, ok := lfu.elements[key]
	if !ok {
		return nil
	}
	node := e.Value.(node)
	lfu.freqList[node.freq].Remove(e)
	node.freq++
	lfu.insertIntoFreq(node)
	return node.value
}

func (lfu LFUCache) Put(key int, value interface{}) {
	if e, ok := lfu.elements[key]; ok {
		node := e.Value.(node)
		lfu.freqList[node.freq].Remove(e)
		node.freq++
		node.value = value

		lfu.insertIntoFreq(node)
		return
	}

	if lfu.Len() < lfu.capacity {
		lfu.insertIntoFreq(node{
			key:   key,
			value: value,
			freq:  1,
		})
		lfu.capacity++
		return
	}

	if lfu.Len() == lfu.capacity {
		l, ok := lfu.freqList[lfu.minFreq]
		for !ok || l.Len() == 0 {
			lfu.minFreq++
			l, ok = lfu.freqList[lfu.minFreq]
		}
		l.Remove(l.Back())
		lfu.insertIntoFreq(node{
			key:   key,
			value: value,
			freq:  1,
		})
		lfu.minFreq = 1
	}
}

func (lfu LFUCache) insertIntoFreq(n node) {
	var e *list.Element
	if l, ok := lfu.freqList[n.freq]; ok {
		e = l.PushFront(n)
		lfu.elements[n.key] = e
	} else {
		l = list.New()
		lfu.freqList[n.freq] = l
		e = l.PushFront(n)
		lfu.elements[n.key] = e
	}
}

func (lfu LFUCache) Len() int {
	return len(lfu.elements)
}

func (lfu LFUCache) contains(key int) {

}

type MapCache map[int]interface{}

func (m MapCache) Put(key int, value interface{}) {
	m[key] = value
}

func (m MapCache) Get(key int) interface{} {
	return m[key]
}

func (m MapCache) Len() int {
	return len(m)
}
