package util

import (
	"fmt"
	"sort"
	"testing"
)

func Test_min(t *testing.T) {
	a := [][]int{{5, 6}, {4, 7}, {4, 8}}
	var r Task = a
	sort.Sort(r)
}

type Task [][]int

func (t Task) Len() int {
	return len(t)
}

func (t Task) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Task) Less(i, j int) bool {
	if t[i][0] == t[j][0] {
		return t[j][1] < t[i][1]
	}
	return t[i][0] < t[j][0]
}

func Test_leetcode(t *testing.T) {
	indexer := map[byte]int{
		'c': 0, 'r': 1, 'o': 2, 'a': 3, 'k': 4,
	}
	fmt.Println(indexer['5'])
}
