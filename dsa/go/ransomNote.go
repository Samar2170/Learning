package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	cmap := make(map[rune]int)

	for _, m := range magazine {
		if _, ok := cmap[m]; !ok {
			cmap[m] = 1
		} else {
			cmap[m]++
		}
	}
	for _, r := range ransomNote {
		if _, ok := cmap[r]; !ok {
			return false
		} else {
			if cmap[r] == 0 {
				return false
			} else {
				cmap[r]--
			}
		}
	}
	return true
}

func main() {
	rn, mag := "aaa", "ba"
	fmt.Println(canConstruct(rn, mag))
	rn, mag = "aa", "ab"
	fmt.Println(canConstruct(rn, mag))
	rn, mag = "aa", "aab"
	fmt.Println(canConstruct(rn, mag))

}
