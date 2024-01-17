package main

import "fmt"

func maxArea(height []int) int {
	n := len(height)
	i, j := 0, n-1
	maxA := 0
	for i < j {
		areaHere := findMin(height[i], height[j]) * (j - i)
		if areaHere > maxA {
			maxA = areaHere
		}
		if height[j] > height[i] {
			i++
		} else {
			j--
		}
	}
	return maxA
}

func findMin(l, r int) int {
	if l < r {
		return l
	} else {
		return r
	}
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
