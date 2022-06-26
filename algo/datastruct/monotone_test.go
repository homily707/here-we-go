package datastruct

import (
	"testing"
)

type Int struct {
	val int
}

func Test_monotone(t *testing.T) {
	deque := MonotoneDeque{comparator: func(a, b interface{}) int {
		return a.([]int)[1] - b.([]int)[1]
	}}
	deque.Push([]int{1, 10})
	deque.Println()
	deque.Push([]int{1, 11})
	deque.Println()
	deque.Push([]int{1, 8})
	deque.Println()
	deque.Push([]int{1, 9})
	deque.Println()

	// LEETCODE 239
	//deque := MonotoneDeque{comparator:func(a,b interface{}) int {
	//	return a.([]int)[1] - b.([]int)[1]
	//}}
	//ans := []int{}
	//
	//for i:=0; i<k-1; i++ {
	//	deque.Push([]int{i,nums[i]})
	//}
	//
	//for i:=k-1; i < len(nums); i++ {
	//	deque.Push([]int{i,nums[i]})
	//	for deque.First().([]int)[0] <= i-k {
	//		deque.PopFirst()
	//	}
	//	ans = append(ans, deque.First().([]int)[1])
	//}
	//
	//return ans
}
