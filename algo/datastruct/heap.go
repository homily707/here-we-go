package datastruct

import (
	"container/heap"
	"fmt"
)

type Ints struct {
	array *[]int
}

func (ints Ints) Len() int {
	return len(*ints.array)
}

func (ints Ints) Less(i, j int) bool {
	return (*ints.array)[i] < (*ints.array)[j]
}

func (ints Ints) Swap(i, j int) {
	(*ints.array)[i], (*ints.array)[j] = (*ints.array)[j], (*ints.array)[i]
}

func (ints Ints) Push(x any) {
	(*ints.array) = append((*ints.array), x.(int))
}

func (ints Ints) Pop() any {
	min := (*ints.array)[len((*ints.array))-1]
	*ints.array = (*ints.array)[:len((*ints.array))-1]
	return min
}

func priorityQueue() {
	ints := Ints{&[]int{4, 5, 6, 7, 1, 2, 3, 4, 5, 8, 9, 10}}
	heap.Init(ints)
	fmt.Println((*ints.array))
	fmt.Println("pop", heap.Pop(ints))
	fmt.Println((*ints.array))
	fmt.Println("pop", heap.Pop(ints))
	fmt.Println((*ints.array))
	heap.Push(ints, 0)
	fmt.Println((*ints.array))
	fmt.Println("pop", heap.Pop(ints))
	fmt.Println((*ints.array))
}
