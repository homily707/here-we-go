package main

import (
	"bytes"
	"sort"
)

func example() {
	abcs := bytes.Repeat([]byte{'a', 'b', 'c'}, 10)

	sort.Ints([]int{1, 3, 4, 5, 6, 2, 1, 3, 4})
	sort.SearchInts([]int{1, 2}, 10)

	print(string(abcs))
}

func main() {
	example()
}
