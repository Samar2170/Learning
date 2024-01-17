package main

import "fmt"

func removeDuplicates(nums []int) int {
	cmap := make(map[int]bool)
	expectedNums := []int{}
	for _, i := range nums {
		if _, ok := cmap[i]; ok {
			continue
		} else {
			cmap[i] = true
			expectedNums = append(expectedNums, i)
		}
	}
	nums = expectedNums
	return len(nums)
}

func removeDuplicates2(nums []int) int {
	prev, i := len(nums)-1, 0
	l := 0
	for i < len(nums)-2 {
		l = i + 1
		if nums[l] == nums[i] {
			nums[prev], nums[l] = nums[l], nums[prev]
			prev--
		}
		i++
	}
	return len(nums)
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates2(nums))
}
