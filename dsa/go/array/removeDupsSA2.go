package main

import "fmt"

func removeDuplicates(nums []int) int {
	w, c, n := 0, 0, nums[0]

	for r := range nums {
		if nums[r] != n || (nums[r] == n && c < 2) {
			if nums[r] != n {
				c = 0
				n = nums[r]
			}
			nums[w] = n
			w++
			c++
		}
	}
	fmt.Println(nums[:w])
	return w
}

func main() {
	// n1 := []int{1, 1, 1, 2, 2, 3}
	// removeDuplicates(n1)

	n2 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	removeDuplicates(n2)

	n3 := []int{1, 1, 1, 1}
	removeDuplicates(n3)

	n4 := []int{0, 0, 0, 0, 0}
	removeDuplicates(n4)
}
