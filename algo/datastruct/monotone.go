package datastruct

import (
	"fmt"
	"reflect"
)

type MonotoneDeque struct {
	data       []interface{}
	comparator func(a, b interface{}) int
}

func (m MonotoneDeque) Println() {
	for _, p := range m.data {
		fmt.Print(reflect.ValueOf(p), " ")
	}
	fmt.Println()
}

func (m *MonotoneDeque) Push(v interface{}) {
	for m.Last() != nil {
		if m.comparator(v, m.Last()) < 0 {
			m.data = append(m.data, v)
			return
		} else {
			m.PopLast()
		}
	}
	m.data = append(m.data, v)
}

func (m MonotoneDeque) Last() interface{} {
	if len(m.data) == 0 {
		return nil
	}
	return m.data[len(m.data)-1]
}

func (m MonotoneDeque) First() interface{} {
	if len(m.data) == 0 {
		return nil
	}
	return m.data[0]
}

func (m *MonotoneDeque) PopLast() interface{} {
	last := m.data[len(m.data)-1]
	m.data = m.data[0 : len(m.data)-1]
	return last
}

func (m *MonotoneDeque) PopFirst() interface{} {
	if len(m.data) == 0 {
		return nil
	}
	first := m.data[0]
	m.data = m.data[1:]
	return first
}
