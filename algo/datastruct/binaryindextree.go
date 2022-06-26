package datastruct

// Integer
type BinaryIndexTree struct {
	Data     []int
	SumSlice []int
	Len      int
}

func (bit *BinaryIndexTree) Init(array []int) {
	bit.Len = len(array)
	bit.Data = array
	bit.SumSlice = make([]int, bit.Len+1)
	for i, num := range array {
		bit.addInSum(i, num)
	}
}

func (bit *BinaryIndexTree) Add(index int, delta int) {
	bit.Data[index] += delta
	bit.addInSum(index, delta)
}

func (bit *BinaryIndexTree) Sum(left int, right int) int {
	return bit.getPreSum(right-1) - bit.getPreSum(left-1)
}

func (bit *BinaryIndexTree) addInSum(index int, delta int) {
	// this is necessary, the real data index starts from 1
	index++
	for index <= bit.Len {
		bit.SumSlice[index] += delta
		index += lowbit(index)
	}
}

func (bit BinaryIndexTree) getPreSum(pos int) int {
	index := pos + 1
	sum := 0
	for index > 0 {
		sum += bit.SumSlice[index]
		index -= lowbit(index)
	}
	return sum
}

func lowbit(x int) int {
	return x & (-x)
}
