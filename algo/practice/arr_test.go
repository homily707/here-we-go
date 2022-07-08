package practice

import (
	"fmt"
	"sort"
	"testing"
)

func TestName(t *testing.T) {
	a := []int{3, 8, -10, 23, 19, -4, -14, 27}
	arr := Abs(a)
	arr.Swap(0, 4)
	fmt.Println(arr)
	sort.Sort(arr)
	fmt.Println(a)
}

func TestString(t *testing.T) {
}

type Abs []int

func (abs Abs) Len() int {
	return len(abs)
}

func (abs Abs) Swap(i, j int) {
	abs[i], abs[j] = abs[j], abs[i]
}

func (abs Abs) Less(i, j int) bool {
	return absolute(abs[i]) < absolute(abs[j])
}

func absolute(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
