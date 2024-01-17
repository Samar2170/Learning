package main

import "fmt"

func diagonalDifference(arr [][]int32) int32 {
	var leftD, rightD int32 = 0, 0
	i := 0
	for i < len(arr) {
		leftD += arr[i][i]
		rightD += arr[i][len(arr[i])-i-1]
		i++
	}
	return abs(leftD - rightD)
}

func abs(i int32) int32 {
	if i < 0 {
		return i * -1
	}
	return i
}

func main() {
	// arr := [][]int32{
	// 	{3},
	// 	{11, 2, 4},
	// 	{4, 5, 6},
	// 	{10, 8, -12},
	// }
	arr2 := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	// fmt.Println(diagonalDifference(arr))
	fmt.Println(diagonalDifference(arr2))
}
