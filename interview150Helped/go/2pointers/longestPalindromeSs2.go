package main

import (
	"fmt"
	"strings"
)

func longestPalindrome(s string) string {
	T := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"
	n := len(T)
	P := make([]int, n)
	C, R := 0, 0
	var x, y int
	for i := 1; i < n-1; i++ {
		fmt.Println(P)
		fmt.Println(C, R)
		fmt.Println(i)
		if R > i {
			x = P[2*C-i]
			y = R - i
			P[i] = min(x, y)
		}
		for T[i+1+P[i]] == T[i-1-P[i]] {
			P[i]++
		}
		if i+P[i] > R {
			C, R = i, i+P[i]
		}
	}
	maxLen := 0
	centerIndex := 0
	for i, v := range P {
		if v > maxLen {
			maxLen = v
			centerIndex = i
		}
	}
	return s[(centerIndex-maxLen)/2 : (centerIndex+maxLen)/2]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	fmt.Println(longestPalindrome("abaxabaxabb"))
}
