package main

func lengthOfLongestSubstring(input string) int {
	// l,r pointers, maxlen to keep track of the longest substring
	// charMap to keep track of the characters in the substring
	l, r := 0, 0
	maxLen := 0
	charMap := make(map[byte]int)

	for r < len(input) {
		// if the character at r is in the map, delete the character at l, increment l
		// else add the character at r to the map and increment r
		if _, ok := charMap[input[r]]; ok {
			delete(charMap, input[l])
			l++
		} else {
			charMap[input[r]] = 1
			r++
		}
		// if the length of the substring is greater than maxLen, update maxLen
		cs := input[l:r]
		if len(cs) > maxLen {
			maxLen = len(cs)
		}
	}
	return maxLen
}

func main() {
	println(lengthOfLongestSubstring("abcabcbb"))
	println(lengthOfLongestSubstring("bbbbb"))
	println(lengthOfLongestSubstring("pwwkew"))
}
