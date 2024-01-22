package algorithm

import (
	"sort"
	"testing"
)

func Test2981(t *testing.T) {
	t.Log(maximumLength("aaabaaabaaaa"))
}

func maximumLength(s string) int {
	count := make(map[rune][]int)
	head := 0
	last := ' '
	for i,r := range s {
			if r == last {
				 if i == len(s)-1 {
					  if arr,ok := count[r]; !ok {
						    count[r] = []int{i-head+1}
				    } else {
						    count[r] = append(arr, i-head+1)
				    }
				 }
					continue
			} else {
					if arr,ok := count[last]; !ok {
							count[last] = []int{i-head}
					} else {
							count[last] = append(arr, i-head)
					}
					head = i
					last = r
					if i == len(s) - 1 {
					  if arr,ok := count[r]; !ok {
						    count[r] = []int{1}
				    } else {
						    count[r] = append(arr, 1)
				    }

					}
			}
	}
	
	ans := -1
	for _,v := range count {
			sort.Slice(v, func(i,j int) bool {
					return i < j
			})
			for i := range v {
					if i+2 < len(v) && v[i+2] == v[i] {
							ans = max(ans, v[i])
							break
					}
					if i+1 < len(v) && v[i+1] == v[i]-1 {
							ans = max(ans, v[i]-1)
							break
					}
					if v[i] >= 3 {
							ans = max(ans, v[i]-2)
					} 
			}
	}
	return ans
}

func max(i,j int) int {
	if i > j {
			return i
	}
	return j
}