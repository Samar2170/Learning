package main

import "fmt"

func quickSort(arr []int, low, high int) []int {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
		fmt.Println(arr)

	}
	fmt.Println("i", i)
	arr[i], arr[high] = arr[high], arr[i]
	fmt.Println(arr)
	return i
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println(quickSort(arr, 0, len(arr)-1))
}
