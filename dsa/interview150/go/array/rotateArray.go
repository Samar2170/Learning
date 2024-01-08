package main

func rotate(nums []int, k int) {
	size := len(nums)
	k = k % size
	c, s := 0, 0
	for c < s {
		ni := s
		temp := nums[ni]
		for {
			ni = (ni + k) % size
			nums[ni], temp = temp, nums[ni]
			c++
			if ni == s {
				break
			}
		}
		s++
	}
}
