package maker

import (
	"math/rand"
)

func NumberArrayInOrder(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

func NumberArrayRandomNotRepeat(n int) []int {
	originArray := NumberArrayInOrder(n)
	swapFunc := func(x, y int) {
		originArray[x], originArray[y] = originArray[y], originArray[x]
	}
	rand.Shuffle(n, swapFunc)
	return originArray
}
