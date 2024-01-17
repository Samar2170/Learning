package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	mergedArray := []int{}
	for i < len(nums1) || j < len(nums2) {
		if i < len(nums1) && j < len(nums2) {
			min := findMin(nums1[i], nums2[j])
			if min == 0 {
				mergedArray = append(mergedArray, nums1[i])
				i += 1
			} else {
				mergedArray = append(mergedArray, nums2[j])
				j += 1
			}
		} else if i < len(nums1) {
			mergedArray = append(mergedArray, nums1[i])
			i++
		} else {
			mergedArray = append(mergedArray, nums2[j])
			j++
		}
	}
	if len(mergedArray)%2 == 0 {
		return float64(mergedArray[(len(mergedArray)/2)-1]+mergedArray[len(mergedArray)/2]) / 2
	} else {
		return float64(mergedArray[(len(mergedArray) / 2)])
	}

	// newArray := []int{}
	// for i < max/2 {
	// 	sortedNums := sortNum(nums1[i], nums2[i])
	// 	newArray = append(newArray, sortedNums...)
	// 	i += 2
	// }
	// return float64(newArray[len(newArray)-1])
}
func findMin(l, r int) int {
	if l < r {
		return 0
	}
	return 1
}

func sortNum(l, r int) []int {
	if l < r {
		return []int{l, r}
	}
	return []int{r, l}
}

func findMax(l, r int) int {
	if l > r {
		return l
	}
	return r
}
func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	ans := findMedianSortedArrays(nums1, nums2)
	fmt.Println(ans)
}
