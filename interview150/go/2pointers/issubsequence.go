package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	i, j := 0, 0
	smatch := make(map[byte]int)
	for i < len(t) {
		if t[i] == s[j] {
			smatch[s[j]] = i
			j++
		}
		i++
	}
	fmt.Println(smatch)
	return j == len(s)

}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc")) // true
	fmt.Println(isSubsequence("axc", "ahbgdc")) // false
}
