package main

func longestCommonPrefix(strs []string) string {
	min := strs[0]

	// assign smallest string to min
	for i := 0; i < len(strs); i++ {
		min = findMinLen(min, strs[i])
	}
	minLen := len(min)
	// start a loop from minLen to 0
	for minLen > 0 {
		// get the first minLen characters of the first string
		currentMin := strs[0][0:minLen]
		matches := 0
		// loop through all the strings and compare the first minLen characters
		for _, s := range strs {
			cs := s[0:minLen]
			// if all the strings have the same first minLen characters, increment matches
			// else break the loop
			if cs == currentMin {
				matches++
				continue
			} else {
				break
			}
		}
		// if matches is equal to the length of the strings, break the loop
		if matches == len(strs) {
			break
		} else {
			// else decrement minLen
			minLen--
		}
	}
	// return the first minLen characters of the first string
	return strs[0][0:minLen]
}
func findMinLen(l, r string) string {
	if len(l) < len(r) {
		return l
	}
	return r
}

func longestCommonPrefix2(strs []string) string {
	min, minlen := strs[0], len(strs[0])
	for _, s := range strs {
		if len(s) < minlen {
			min = s
			minlen = len(s)
		}
	}
	return min
}
func main() {
	strs := []string{"flower", "flow", "flight"}
	println(longestCommonPrefix(strs))
	strs = []string{"dog", "racecar", "car"}
	println(longestCommonPrefix(strs))

	strs = []string{"flower", "flow", "flight"}
	println(longestCommonPrefix2(strs))
	strs = []string{"reflower", "flow", "flight"}
	println(longestCommonPrefix2(strs))
}
