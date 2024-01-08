package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	i := 1
	for i < len(nums) {
		if i == len(nums)-1 && i != 1 {
			return true
		} else if i > len(nums)-1 || nums[i] == 0 {
			return false
		}
		i = i + nums[i]
	}
	return false
}

func main() {
	nums1 := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums1))
	nums2 := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums2))
	nums3 := []int{0, 1}
	fmt.Println(canJump(nums3))
	nums4 := []int{1, 2}
	fmt.Println(canJump(nums4))
}
