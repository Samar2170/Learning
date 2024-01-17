package main

import (
	"fmt"
	"strings"
)

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	i := 0
	longestPd := ""
	for i < len(s)-1 {
		pd := expandFromCentre(s, i)
		if len(pd) > len(longestPd) {
			longestPd = pd
		}
		i++
	}
	return longestPd
}
func expandFromCentre(st string, i int) string {
	var j, k int
	j, k = i-1, i+1
	pd := ""
	builder := &strings.Builder{}
	if len(st)%2 == 0 {
		res := []string{"$"}
		for _, s := range st {
			res = append(res, string(s))
			res = append(res, "$")
		}
		for _, re := range res {
			builder.WriteString(re)
		}
		st = builder.String()
	}
	for j > 0 && k < len(st) {
		if st[j] == st[k] {
			pd = st[j : k+1]
			k++
			j--
		} else {
			break
		}
	}
	return pd
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("ac"))
	fmt.Println(longestPalindrome("bb"))
	fmt.Println(longestPalindrome("ccc"))
	fmt.Println(longestPalindrome("aaaa"))
	fmt.Println(longestPalindrome("aaaaa"))
	fmt.Println(longestPalindrome("abaxabaxabb"))
}
