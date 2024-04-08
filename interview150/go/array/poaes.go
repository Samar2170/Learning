package main

import "fmt"

func productExceptSelf(nums []int) []int {
	pt := 1
	for _, n := range nums {
		pt = pt * n
	}
	newArray := []int{}
	for _, n := range nums {
		newArray = append(newArray, pt/n)
	}
	return newArray
}

func productExceptSelf2(nums []int) []int {
	n := len(nums)
	prefix, suffix := []int{}, []int{}
	for _, _ = range nums {
		prefix = append(prefix, 1)
		suffix = append(suffix, 1)
	}
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}
	output := []int{}
	for i := 0; i < n; i++ {
		output = append(output, prefix[i]*suffix[i])
	}
	return output
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
	fmt.Println(productExceptSelf2(nums))
}
