package datastruct

//
//import (
//	"fmt"
//	"reflect"
//)
//
//type GeMonotoneDeque [T any] struct {
//	data []T
//	comparator func(a interface{},b interface{}) int
//}
//
//func (m *GeMonotoneDeque[T]) With(f func(a interface{},b interface{}) int) {
//	m.comparator = f
//}
//
//func (m GeMonotoneDeque[T]) Println() {
//	for _, p := range m.data {
//		fmt.Print(reflect.ValueOf(p), " ")
//	}
//	fmt.Println()
//}
//
//func (m *GeMonotoneDeque[T]) Push(v T) {
//	for m.Last() != *new(T) {
//		if m.comparator(v, m.Last()) > 0{
//			m.data = append(m.data, v)
//			return
//		} else {
//			m.PopLast()
//		}
//	}
//	m.data = append(m.data, v)
//}
//
//func (m GeMonotoneDeque[T]) Last() T {
//	if len(m.data) == 0 {
//		return *new(T)
//	}
//	return m.data[len(m.data)-1]
//}
//
//func (m GeMonotoneDeque[T]) First() T {
//	if len(m.data) == 0 {
//		return *new(T)
//	}
//	return m.data[0]
//}
//
//func (m *GeMonotoneDeque[T]) PopLast() T {
//	last := m.data[len(m.data)-1]
//	m.data = m.data[0 : len(m.data)-1]
//	return last
//}
//
//func (m *GeMonotoneDeque[T]) PopFirst() T {
//	if len(m.data) == 0 {
//		return *new(T)
//	}
//	first := m.data[0]
//	m.data = m.data[1:]
//	return first
//}
