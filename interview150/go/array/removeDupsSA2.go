package main

import "fmt"

func removeDuplicates(nums []int) int {
	i := 0
	numMap := make(map[int]struct{})
	for i < len(nums) {
		if _, ok := numMap[nums[i]]; ok {
			oldArr := nums
			nums = nums[:i]
			nums = append(nums, oldArr[i+1:]...)
			i--
		} else {
			numMap[nums[i]] = struct{}{}
		}
		i++
	}
	return len(nums)
}

func main() {
	nums := []int{1, 1, 2}
	removeDuplicates(nums)
	fmt.Println(nums)

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(nums)
	fmt.Println(nums)
}
