package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	nmap := make(map[int]int)
	for i, n := range numbers {
		if _, ok := nmap[n]; ok {
			return []int{nmap[n] + 1, i + 1}
		} else {
			nmap[target-n] = i
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0, 1]
	fmt.Println(twoSum([]int{2, 3, 4}, 6))      // [0, 2]
	fmt.Println(twoSum([]int{-1, 0}, -1))       // [0, 1]
}
