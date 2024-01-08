package main

import "fmt"

func main() {
	var s string
	// strs := []string{"flower", "flow", "flight"}
	strs2 := []string{"caa", "", "a", "acb"}
	// s = longestCommonPrefix(strs)
	s = longestCommonPrefix(strs2)
	fmt.Println(s)

}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	var min string

	i := 0
	min = strs[i]
	for i < len(strs)-1 {
		min = findMinLen(min, strs[i+1])
		i++
	}
	minLen := len(min)
	for minLen > 0 {
		mins := strs[0][0:minLen]
		matches := 0
		for _, s := range strs {
			if minLen == 0 {
				break
			}
			cs := s[0:minLen]
			if cs == mins {
				matches++
				continue
			} else {
				break
			}
		}
		if matches == len(strs) {
			break
		} else {
			minLen--
		}
	}
	return strs[0][0:minLen]
}

func findMinLen(l, r string) string {
	if len(l) < len(r) {
		return l
	}
	return r
}
