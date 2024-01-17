package main

import "fmt"

func majorityElement(nums []int) int {
	nmap := make(map[int]int)
	for _, n := range nums {
		if _, ok := nmap[n]; ok {
			nmap[n]++
		} else {
			nmap[n] = 1
		}
	}
	for k, v := range nmap {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}

func main() {
	n1 := []int{2, 2, 1, 1, 1, 2, 2}
	n2 := []int{3, 2, 3}
	n3 := []int{1}
	fmt.Println(majorityElement(n1))
	fmt.Println(majorityElement(n2))
	fmt.Println(majorityElement(n3))

}
