package main

import "fmt"

func removeElement(nums []int, val int) int {
	remove := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			remove++
			i--
		}
	}
	return len(nums) - remove
}

// Path: interview150/go/array/rotate.go

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	removeElement(nums, 3)
	fmt.Println(nums)

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	removeElement(nums, 2)
	fmt.Println(nums)
}
