package main

import "fmt"

func isValid(s string) bool {
	stack := []string{}

	for _, sc := range s {
		char := string(sc)
		fmt.Println(char, stack)
		switch char {
		case "(", "[", "{":
			stack = append(stack, char)
		case ")":
			if len(stack) > 0 && stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case "}":
			if len(stack) > 0 && stack[len(stack)-1] == "{" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case "]":
			if len(stack) > 0 && stack[len(stack)-1] == "[" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println("Valid Parenthesis")
	s := "()[]{}"
	fmt.Println(isValid(s))
	s = "([)]"
	fmt.Println(isValid(s))
	s = "{[]}"
	fmt.Println(isValid(s))
	s = "){"
	fmt.Println(isValid(s))
}
