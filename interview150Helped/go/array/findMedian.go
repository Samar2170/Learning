package main

import "fmt"

func findMedian(arr []int32) int32 {
	// Write your code here
	sortedArray := bubbleSort(arr)
	fmt.Println(arr)
	fmt.Println(sortedArray)
	return sortedArray[len(arr)/2]
}
func bubbleSort(arr []int32) []int32 {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}

func main() {
	ints := []int32{
		0, 1, 2, 4, 6, 5, 3,
	}
	fmt.Println(findMedian(ints))

	fmt.Println(bubbleSort(ints))

	//  testcase to break thi

}
