package main

import "fmt"

func isValid(s string) bool {
	stack := []string{}
	elMap := map[string]struct{}{
		"[": {},
		"{": {},
		"(": {},
	}
	for _, el := range s {
		els := string(el)
		if _, ok := elMap[els]; ok {
			stack = append(stack, els)
		}
		if els == "]" {
			if stack[len(stack)-1] != "[" {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
		if els == "}" {
			if stack[len(stack)-1] != "{" {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
		if els == ")" {
			if stack[len(stack)-1] != "(" {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func isValidWithAsterisk(s string) bool {
	stack := []string{}
	elMap := map[string]struct{}{
		"[": {},
		"{": {},
		"(": {},
	}
	for _, el := range s {
		els := string(el)
		if _, ok := elMap[els]; ok {
			stack = append(stack, els)
		}
		if els == "]" {
			if stack[len(stack)-1] == "[" || stack[len(stack)-1] == "*" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
		if els == "}" {
			if stack[len(stack)-1] == "{" || stack[len(stack)-1] == "*" {
				stack = stack[:len(stack)-1]

			} else {
				return false
			}
		}
		if els == ")" {
			if stack[len(stack)-1] == "(" || stack[len(stack)-1] == "*" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
		if els == "*" {
			fmt.Println(stack)
			if _, ok := elMap[stack[len(stack)-1]]; ok {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, els)
			}

		}
	}
	return len(stack) == 0
}

func main() {
	println(isValid("()"))
	println(isValid("()[]{}"))
	println(isValid("(]"))
	println(isValid("([)]"))
	println(isValid("{[]}"))
	println(isValidWithAsterisk("[(*){*)]"))
}
