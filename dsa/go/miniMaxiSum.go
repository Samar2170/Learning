package main

import "fmt"

func miniMaxSum(arr []int32) {
	// Write your code here
	sortedArray := bubbleSort(arr)
	var miniSum, maxiSum int64 = 0, 0
	for indx, n := range sortedArray {
		fmt.Println(miniSum, maxiSum)
		if indx == 0 {
			miniSum += int64(n)
		} else if indx == len(arr)-1 {
			maxiSum += int64(n)
		} else {
			miniSum += int64(n)
			maxiSum += int64(n)
		}
	}
	fmt.Println(miniSum, maxiSum)
}
func bubbleSort(arr []int32) []int32 {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	ints := []int32{
		396285104, 573261094, 759641832, 819230764, 364801279,
	}
	miniMaxSum(ints)

	fmt.Println(bubbleSort(ints))

}
