package main

import "fmt"

// self
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
	fmt.Println(nums)
	return len(nums)
}

func main() {
	n1 := []int{1, 1, 2}
	n2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(n1)
	removeDuplicates(n2)
}
