package main

import (
	"fmt"
)

func removeDuplicates2(nums []int) int {
	writer, counter, numberTracker := 0, 0, nums[0]

	for r := range nums {
		if (nums[r] != numberTracker) || (nums[r] == numberTracker && counter < 2) {
			if nums[r] != numberTracker {
				numberTracker = nums[r]
				counter = 0
			}
			nums[writer] = nums[r]
			writer++
			counter++
		}
	}
	return len(nums[:writer])
}
func removeDuplicates(nums []int) int {
	nmap := make(map[int]int)
	for _, n := range nums {
		if _, ok := nmap[n]; ok {
			nmap[n]++
		} else {
			nmap[n] = 1
		}
	}

	for i, n := range nums {
		if nmap[n] > 2 {
			if i == len(nums)-1 {
				nums[i] = -9999
				nmap[n]--
			} else {
				nums[i], nums[i+1] = nums[i+1], -1
				nmap[n]--
			}
		}
	}
	i := 0
	for i < len(nums) {
		if nums[i] == -9999 {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}

	}
	fmt.Println(nums)
	return len(nums)
}

func main() {
	n1 := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates2(n1))

	n2 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates2(n2))

	n3 := []int{1, 1, 1, 1}
	fmt.Println(removeDuplicates2(n3))

	n4 := []int{0, 0, 0, 0, 0}
	fmt.Println(removeDuplicates2(n4))
}
