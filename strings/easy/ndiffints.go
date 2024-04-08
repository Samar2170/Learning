package main

import (
	"fmt"
)

func numDifferentIntegers(word string) int {
	m := make(map[string]bool)
	i := 0
	for i < len(word) {
		if word[i] >= '0' && word[i] <= '9' {
			j := i
			for j < len(word) && word[j] >= '0' && word[j] <= '9' {
				j++
			}
			for i < j && word[i] == '0' {
				i++
			}
			m[word[i:j]] = true
			i = j
		} else {
			i++
		}
	}
	return len(m)
}

func main() {
	fmt.Println(numDifferentIntegers("a123bc34d8ef34"))  // 3
	fmt.Println(numDifferentIntegers("leet1234code234")) // 2
	fmt.Println(numDifferentIntegers("a1b01c001"))       // 1
}
