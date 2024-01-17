package main

import "fmt"

func rotate(nums []int, k int) {
	i := 0
	n := len(nums)
	for i < k {
		re := nums[0 : n-1]
		le := nums[n-1]
		nums = []int{}
		nums = append(nums, le)
		nums = append(nums, re...)
		i++
	}
	fmt.Println(nums)

}

func main() {
	n1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(n1, 3)

	n2 := []int{-1, -100, 3, 99}
	rotate(n2, 2)

	n3 := []int{1, 2}
	rotate(n3, 3)
}
