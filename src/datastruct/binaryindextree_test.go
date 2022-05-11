package datastruct

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBinaryIndexTree(t *testing.T) {
	array := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	var bit BinaryIndexTree = BinaryIndexTree{}
	bit.Init(array)
	fmt.Println(bit.Data)
	fmt.Println(bit.SumSlice)

	fmt.Print(strconv.Itoa(bit.Sum(6, 9)) + " | ")
	bit.Add(9, 10)
	bit.Add(7, 10)

	fmt.Println(bit.Sum(6, 9))

	fmt.Println(bit.Data)
	fmt.Println(bit.SumSlice)
}
