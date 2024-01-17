package main

import "fmt"

func rotateNew(nums []int, k int) {
	newArray := []int{}

	lastEl := nums[len(nums)-k:]
	firstEl := nums[:len(nums)-k]
	newArray = append(newArray, lastEl...)
	newArray = append(newArray, firstEl...)

	fmt.Println(newArray)

}

func rotate(nums []int, k int) {
	size := len(nums)
	k = k % size

	var count, start = 0, 0

	for count < size {
		nextIdx := start
		temp := nums[start]
		for {
			nextIdx = (nextIdx + k) % size
			temp, nums[nextIdx] = nums[nextIdx], temp
			count++
			if start == nextIdx {
				break
			}
		}
		start++
	}
	fmt.Println(nums)
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)

	rotate([]int{-1, -100, 3, 99}, 2)

}
