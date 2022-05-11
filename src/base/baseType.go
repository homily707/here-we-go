package main

import (
	"fmt"
)

func main() {
	//3 way to claim a var, take array for example
	// var array0 [2]int = [2]int {1,2}
	// var array1 = [3]int {1,2,3}
	// array2 := [3]int {}
	// var array3 [3]int = [3]int{}

	//way to create slice
	// slice0 :=  make([]int, 5, 10)
	// slice1 := []int{1,2,3,4,5,6,7} // notice the difference with array1
	// slice2 := []int{}
	// fmt.Println(len(slice1),cap(slice1))

	//array is immutalable
	//append(array2,1)  // first argument to append must be slice; have [3]int
	// slice append must be used
	//append(slice1,4)  // append(slice1, 4) evaluated but not used

	// var s1 = slice1[2:3]

	// map
	m := make(map[int]string, 16)
	//save
	m[0] = "zero"
	fmt.Println(m)
	//read
	str := m[0]
	fmt.Println(m, str)
	//del
	delete(m, 0)
	fmt.Println(m)
	//exist
	_, ok := m[0]
	if ok {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}
}
