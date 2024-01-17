package main

import "fmt"

func mergeSortedArray(nums1, nums2 []int, m, n int) {
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	fmt.Println(nums1)
}

func main() {
	n1 := []int{1, 2, 3, 0, 0, 0}
	n2 := []int{2, 5, 6}
	mergeSortedArray(n1, n2, 3, 3)

	n3 := []int{10, 90, 150, 0, 0, 0}
	n4 := []int{20, 30, 40}
	mergeSortedArray(n3, n4, 3, 3)
}
