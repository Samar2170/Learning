package main

import "fmt"

func insertElements(nums, elems []int) []int {
	for _, el := range elems {
		nums = append(nums, el)
	}
	return nums
}

func deleteElement(nums []int, index int) []int {
	nums = append(nums[:index], nums[index+1:]...)
	return nums
}
func deleteElements(nums []int, indices []int) []int {
	imap := make(map[int]bool)
	for _, i := range indices {
		imap[i] = true
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if _, ok := imap[i]; ok {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return nums
}

func filterArray(nums []int, filterFunc func(int) bool) []int {
	for i, n := range nums {
		if !filterFunc(n) {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return nums
}

func reverseArray(arr []int) {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-1-i] = arr[l-1-i], arr[i]
	}
}

func rotateLeft(arr []int, k int) {
	l := len(arr)
	k = k % l
	reverseArray(arr[:k])
	reverseArray(arr[k:])
	reverseArray(arr)
}

func rotateRight(arr []int, k int) {
	l := len(arr)
	k = k % l
	reverseArray(arr[l-k:])
	reverseArray(arr[:l-k])
	reverseArray(arr)
}

func main() {
	index := 2
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(deleteElement(nums, index))

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	indices := []int{2, 4, 5, 6}
	fmt.Println(deleteElements(nums, indices))

	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ff := func(n int) bool {
		return n%2 == 0
	}
	fmt.Println(filterArray(nums, ff))

	nums = []int{1, 2, 3, 4, 5}
	reverseArray(nums)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 4, 5}
	rotateLeft(nums, 2)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 4, 5}
	rotateRight(nums, 2)
	fmt.Println(nums)

}
