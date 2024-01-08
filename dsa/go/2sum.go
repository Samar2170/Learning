package main

func twoSum(nums []int, target int) []int {
	kv := make(map[int]int)
	for indx, x := range nums {
		if _, ok := kv[x]; ok {
			return []int{
				indx, kv[x],
			}
		}
		kv[target-x] = indx
	}
	return []int{}
}
