package main

func lengthOfLongestSubstring(s string) int {
	i, j := 0, 0
	maxlen := 0
	charMap := make(map[string]uint)
	for j < len(s) {
		r := string(s[j])
		charMap[r] += 1
		for charMap[r] > 1 {
			l := string(s[i])
			charMap[l] -= 1
			i += 1
		}
		maxlen = FindMax(maxlen, j-i+1)
		j += 1
	}
	return maxlen
}

func FindMax(l, r int) int {
	if l > r {
		return l
	}
	return r
}
