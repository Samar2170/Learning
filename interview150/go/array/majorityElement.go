package main

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
