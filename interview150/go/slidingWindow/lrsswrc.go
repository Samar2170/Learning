package main

func lengthOfLongestSubstring(input string) int {
	l, r := 0, 0
	maxLen := 0
	charMap := make(map[byte]int)

	for r < len(input) {
		if _, ok := charMap[input[r]]; ok {
			delete(charMap, input[l])
			l++
		} else {
			charMap[input[r]] = 1
			r++
		}
		cs := input[l:r]
		if len(cs) > maxLen {
			maxLen = len(cs)
		}
	}
	return maxLen
}

func main() {
	println(lengthOfLongestSubstring("abcabcbb")) // 3
	println(lengthOfLongestSubstring("bbbbb"))    // 1
	println(lengthOfLongestSubstring("pwwkew"))   // 3
	println(lengthOfLongestSubstring(""))         // 0
	println(lengthOfLongestSubstring(" "))        // 1
	println(lengthOfLongestSubstring("au"))       // 2
	println(lengthOfLongestSubstring("aab"))      // 2
	println(lengthOfLongestSubstring("dvdf"))     // 3
	println(lengthOfLongestSubstring("anviaj"))   // 5
	println(lengthOfLongestSubstring("abba"))     // 2
}
