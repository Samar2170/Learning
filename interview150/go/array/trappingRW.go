package main

import "fmt"

func trap(height []int) int {
	l, r := 0, len(height)-1
	res := 0
	lmax, rmax := 0, 0
	for l < r {
		if height[l] < height[r] {
			if height[l] < lmax {
				res += lmax - height[l]
			} else {
				lmax = height[l]
			}
			l++
		} else {
			if height[r] < rmax {
				res += rmax - height[r]
			} else {
				rmax = height[r]
			}
			r--
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))

	height = []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(height))

}

// Intuition
// Imagine the given array representing a series of blocks where each block's height is given.
// We're essentially finding the amount of water that can be trapped between these blocks.
// By comparing heights from both ends (left and right) and moving towards each other, we simulate how water would accumulate in between blocks with varying heights.
// Whenever we encounter a dip between blocks on either side, we calculate the trapped water amount based on the maximum height encountered so far on that side.
// We keep moving towards the center until we've covered the entire array and calculated the total trapped water.
// Approach
// Setting Up Variables: We start by setting up some variables.

// l starts at the left end of the array (height).
// r starts at the right end of the array.
// res is where we'll store the result, initialized to 0 (no trapped water yet).
// lmax and rmax are initially set to 0, representing the maximum heights encountered from the left and right sides, respectively.
// Looping through the Array: We use a loop to iterate through the array.

// The loop continues until the left and right pointers meet.
// We compare heights at l and r positions.
// Finding Trapped Water: We analyze which side (left or right) has a lower height.

// If the height at l is less than or equal to the height at r, we focus on the left side.
// If the height at l is less than the lmax, it means there's a potential for trapping water. We calculate the trapped water amount by subtracting the current height at l from the lmax. If the current height is greater than lmax, we update lmax to the current height.
// If the height at r is greater than the rmax, we update rmax to the current height. Otherwise, it indicates a potential for trapping water, so we calculate and add the trapped water amount to the result by subtracting the current height at r from the rmax.
// Moving Pointers: We move the pointers (l and r) towards each other.

// If the height at l is less than or equal to the height at r, we move l to the right.
// If the height at l is greater than the height at r, we move r to the left.
// Returning the Result: Finally, we return the total trapped water amount (res)
