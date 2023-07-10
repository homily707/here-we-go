package util

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func slice2d(m, n int) [][]int {
	mn := make([][]int, m)
	for i := range mn {
		mn[i] = make([]int, n)
	}
	return mn
}
