package main

import "fmt"

func longestSubString(input string) string {
	i, j := 0, 1

	maxStr := ""
	for i < len(input) || j < len(input) {
		if input[j] == input[i] {
			i++
			j++
		} else {
			newStr := input[i:j]
			if len(newStr) > len(maxStr) {
				maxStr = newStr

			}
			j++
		}
	}
	return maxStr
}

func main() {
	s := "aaaadevxbc"
	fmt.Println(longestSubString(s))
}
