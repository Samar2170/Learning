package main

import "../../interview150/go/strings/easy/fmt"

// iterate over word
// if word[ith]==sequence[0] then check if substring matches

func maxRepeating(sequence string, word string) int {
	if sequence == word {
		return 1
	}
	l := 0
	slen := len(word)
	matches := 0
	for l < len(sequence)-len(word) {
		if string(sequence[l]) == string(word[0]) {
			if sequence[l:l+slen] == word {
				matches++
			}
		}
		l++
	}
	return matches
}
func maxRepeating2(sequence string, word string) int {
	// iterate over sequence
	// if sequence[i]==word[0], see if word==sequence[i:i+len(word)]
	// move i to i+len(word)
	l, r := 0, 0
	slen := len(word)
	mMatches := 0
	for l < len(sequence)-len(word) {
		matches := 0
		for string(sequence[r]) == string(word[0]) {
			if sequence[r:r+slen] == word {
				matches++
				if r+slen < len(sequence) {
					r = r + slen
				} else {
					break
				}
			} else {
				break
			}
		}
		if matches > mMatches {
			mMatches = matches
		}
		l++
		r = l
	}
	return mMatches
}
func main() {
	fmt.Println(maxRepeating("ababc", "ab"))                                  // 2
	fmt.Println(maxRepeating("ababc", "ba"))                                  // 1
	fmt.Println(maxRepeating("ababc", "ac"))                                  // 0
	fmt.Println(maxRepeating2("aaabaaaabaaabaaaabaaaabaaaabaaaaba", "aaaba")) // 5

}
