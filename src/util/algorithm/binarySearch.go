package main

func main() {
	testArray := []int{1, 2, 3, 4, 6, 7, 8, 9}
	println(testArray)
	println(binarySearch(testArray, 9))
	println(binarySearch(testArray, 5))
}

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l >= len(nums) || nums[l] != target {
		return -1
	} else {
		return l
	}
}
