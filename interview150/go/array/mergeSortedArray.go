package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {

			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)

	fmt.Println(nums1)
	nums1 = []int{1}
	m = 1
	nums2 = []int{}
	n = 0
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

	nums1 = []int{7, 8, 9, 0, 0, 0}
	m = 3
	nums2 = []int{1, 2, 3}
	n = 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

}
