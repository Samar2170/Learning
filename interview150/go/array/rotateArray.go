package main

import "fmt"

func reverseArray(nums []int) {
	l := len(nums)
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[l-i-1] = nums[l-i-1], nums[i]
	}
}

func rotateArray(nums []int, k int) {
	l := len(nums)
	k = k % l
	reverseArray(nums[l-k:])
	reverseArray(nums[:l-k])
	reverseArray(nums)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotateArray(nums, 3)
	fmt.Println(nums)

	nums = []int{-1, -100, 3, 99}
	rotateArray(nums, 2)
	fmt.Println(nums)
}
