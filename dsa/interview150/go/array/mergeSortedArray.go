package main

// helped

func merge(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}
func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1

	for k := m + n - 1; k >= 0; k-- {
		if j < 0 {
			return
		}
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}

}
