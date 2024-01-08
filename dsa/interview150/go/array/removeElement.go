package main

import "fmt"

// self

func removeElement(nums []int, val int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == val {
			oldArr := nums
			nums = oldArr[:i]
			nums = append(nums, oldArr[i+1:]...)
			i--
		}
		i++
	}
	fmt.Println(nums)
	return len(nums)
}

func main() {
	n1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n2 := []int{3, 2, 2, 3}
	n3 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	removeElement(n1, 3)
	removeElement(n2, 3)
	removeElement(n3, 2)
}
