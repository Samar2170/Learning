package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	i := 0

	maxNums := []int{}
	for i < len(nums)-(k-1) {
		j := i + k
		maxHere := nums[i]
		fmt.Println(i, j, k, maxHere)
		for _, n := range nums[i:j] {
			if n > maxHere {
				maxHere = n
			}
		}
		maxNums = append(maxNums, maxHere)
		i++
	}

	return maxNums
}

func main() {
	s1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(s1, 3))

	s2 := []int{1, -1}
	fmt.Println(maxSlidingWindow(s2, 2))

	s3 := []int{9}
	fmt.Println(maxSlidingWindow(s3, 1))

}
