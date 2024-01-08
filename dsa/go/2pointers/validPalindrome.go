package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	pattern, _ := regexp.Compile("[^a-zA-Z0-9]+")
	s = pattern.ReplaceAllString(s, "")
	s = strings.ToLower(s)
	var i, j int
	if len(s)%2 == 1 {
		i, j = (len(s)/2)-1, (len(s)/2)+1
		fmt.Println(i, j, len(s))
	} else {
		i, j = len(s)/2-1, len(s)/2
		fmt.Println(i, j, len(s))
	}
	for i >= 0 && j < len(s) {
		if s[i] != s[j] {
			return false
		}
		i--
		j++
	}
	return true
}
func isPalindrome2(s string) bool {
	pattern, _ := regexp.Compile("[^a-zA-Z0-9]+")
	s = pattern.ReplaceAllString(s, "")
	s = strings.ToLower(s)
	s2 := ""
	for i := len(s) - 1; i >= 0; i-- {
		s2 += string(s[i])
	}
	fmt.Println(s, s2)
	if s == s2 {
		return true
	}
	return false
}
func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome2(s))

	s = "race a car"
	fmt.Println(isPalindrome(s))

	s = "aba"
	fmt.Println(isPalindrome2(s))

	s = "bb"
	fmt.Println(isPalindrome2(s))
	s = "abcdcba" // 7/2-1 = 3, 7/2+1 = 5
	fmt.Println(isPalindrome2(s))
}
