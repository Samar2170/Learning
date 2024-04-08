package main

import "fmt"

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	ma := 0
	for i < j {
		ca := findMin(height[i], height[j]) * (j - i)
		if ca > ma {
			ma = ca
		}
		if i > j {
			j--
		} else {
			i++
		}
	}
	return ma
}

func findMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	fmt.Println(maxArea([]int{1, 1}))                      // 1
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))             // 16
	fmt.Println(maxArea([]int{1, 2, 1}))                   // 2
}
