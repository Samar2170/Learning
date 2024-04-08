package main

func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, op := range operations {
		if op[1] == '+' {
			x++
		} else {
			x--
		}
	}
	return x
}

func main() {
	println(finalValueAfterOperations([]string{"--X", "X++", "X++"}))
	println(finalValueAfterOperations([]string{"++X", "++X", "X++"}))
	println(finalValueAfterOperations([]string{"X++", "++X", "--X", "X--"}))
}
