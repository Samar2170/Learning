package main

// self
func removeDuplicates(nums []int) int {
	i := 0
	numMap := make(map[int]int)
	for i < len(nums) {
		if val, ok := numMap[nums[i]]; ok {
			if val > 1 {
				oldArr := nums
				nums = nums[:i]
				nums = append(nums, oldArr[i+1:]...)
				i--
			} else {
				numMap[nums[i]]++
			}
		} else {
			numMap[nums[i]] = 1
		}
		i++
	}
	// fmt.Println(nums)
	return len(nums)
}

func main() {
	n1 := []int{1, 1, 2}
	n2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n3 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	removeDuplicates(n1)
	removeDuplicates(n2)
	removeDuplicates(n3)
}
