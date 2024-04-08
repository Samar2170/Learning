package main

func minSubArrayLen(target int, nums []int) int {
	i, j := 0, 0
	sum := 0
	minLen := len(nums)
	for i < len(nums) {
		na := []int{}
		for j < len(nums) && sum < target {
			sum += nums[j]
			na = append(na, nums[j])
			j++
		}
		if len(na) < minLen && sum >= target {
			minLen = len(na)
		}
		sum, j = 0, i
		i++
	}
	if minLen == len(nums) {
		csum := 0
		for _, n := range nums {
			csum += n
		}
		if csum < target {
			return 0
		}
	}
	return minLen
}

func main() {
	println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))        // 2
	println(minSubArrayLen(4, []int{1, 4, 4}))                 // 1
	println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) // 0
	println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))          // 3
	println(minSubArrayLen(15, []int{1, 2, 3, 4, 5}))          // 5
	println(minSubArrayLen(7, []int{1, 1, 1, 1, 7}))
}
