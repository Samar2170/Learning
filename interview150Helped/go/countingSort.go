package main

import "fmt"

func countingSort(arr []int32) []int32 {
	countingArray := []int32{}
	var maxNo int32 = 0
	for i, _ := range arr {
		if arr[i] > maxNo {
			maxNo = arr[i]
			if maxNo > int32(len(countingArray)) {
				addL := maxNo - int32(len(countingArray))
				var j int32
				for j < addL {
					countingArray = append(countingArray, 0)
					j++
				}
			}
		}
		countingArray[int(arr[i])-1]++
	}
	return countingArray
}

type Counter map[int32]int32

func (c *Counter) Add(key int32) {
	if _, ok := (*c)[key]; ok {
		(*c)[key]++
	} else {
		(*c)[key] = 1
	}
}

func countingSort2(arr []int32) []int32 {
	var si Counter = make(map[int32]int32)
	ii := make(map[int]int32)
	for i, _ := range arr {
		si.Add(arr[i])
		ii[i] = arr[i]
	}
	// fmt.Println(si)
	// var newArray []int32
	for i, _ := range arr {
		// newArray = append(newArray, si[ii[i]]-1)
		arr[i] = si[ii[i]] - 1
	}
	return arr
}

func main() {
	// arr1 := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// arr2 := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(countingSort(arr1))
	// fmt.Println(countingSort(arr2))

	// arr3 := []int32{63, 25, 73, 1, 98, 73, 56, 84, 86, 57, 16, 83, 8,
	// 	25, 81, 56, 9, 53, 98, 67, 99, 12, 83, 89,
	// 	80, 91, 39, 86, 76, 85, 74, 39, 25, 90,
	// 	59, 10, 94, 32, 44, 3, 89, 30, 27, 79,
	// 	46, 96, 27, 32, 18, 21, 92, 69, 81, 40,
	// }
	// 63 25 73 1 98 73 56 84 86 57 16 83 8 25 81 56 9 53 98 67 99 12 83 89 80 91 39 86 76 85 74 39 25 90 59 10 94 32 44 3 89 30 27 79 46 96 27 32 18 21 92 69 81 40 40 34 68 78 24 87 42 69 23 41 78 22 6 90 99 89 50 30 20 1 43 3 70 95 33 46 44 9 69 48 33 60 65 16 82 67 61 32 21 79 75 75 13 87 70 33
	arr4 := []int32{63, 25, 73, 1, 98, 73, 56, 84, 86, 57, 16, 83, 8, 25, 81, 56, 9, 53, 98, 67, 99, 12, 83, 89, 80, 91, 39, 86, 76, 85, 74, 39, 25, 90, 59, 10, 94, 32, 44, 3, 89, 30, 27, 79, 46, 96, 27, 32, 18, 21, 92, 69, 81, 40, 40, 34, 68, 78, 24, 87, 42, 69, 23, 41, 78, 22, 6, 90, 99, 89, 50, 30, 20, 1, 43, 3, 70, 95, 33, 46, 44, 9, 69, 48, 33, 60, 65, 16, 82, 67, 61, 32, 21, 79, 75, 75, 13, 87, 70, 33}
	// fmt.Println(countingSort2(arr3))
	fmt.Println(arr4)
	fmt.Println(countingSort2(arr4))
}
