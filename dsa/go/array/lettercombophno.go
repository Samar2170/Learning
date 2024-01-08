package main

import "fmt"

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	phoneMap := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	combinations := []string{""}

	for _, digit := range digits {
		newCombinations := []string{}
		for _, comb := range combinations {
			for _, letter := range phoneMap[digit-'2'] {
				newCombinations = append(newCombinations, comb+string(letter))
			}
		}
		combinations = newCombinations
	}
	return combinations
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations("9875"))

}
