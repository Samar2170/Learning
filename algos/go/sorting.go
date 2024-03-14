package main

import "fmt"

func QuickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func MergeSort(arr []int) {
	if len(arr) > 1 {
		mid := len(arr) / 2
		l := arr[:mid]
		r := arr[mid:]
		MergeSort(l)
		MergeSort(r)
		i, j, k := 0, 0, 0
		for i < len(l) && j < len(r) {
			if l[i] <= r[j] {
				arr[k] = l[i]
				i++
			} else {
				arr[k] = r[j]
				j++
			}
			k++
		}
		for i < len(l) {
			arr[k] = l[i]
			i++
			k++
		}
		for j < len(r) {
			arr[k] = r[j]
			j++
			k++
		}
	}
}

func main() {
	nums := []int{4, 3, 2, 1}
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)

	nums = []int{10, 17, 56, 82, 13, 9, 23, 18, 7}
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)

	nums3 := []int{25, 55, 10, 17, 80, 95, 74, 46}
	MergeSort(nums3)
	fmt.Println(nums3)
}
