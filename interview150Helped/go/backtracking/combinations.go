package main

import "fmt"

func main() {
	n, k := 4, 2
	fmt.Println(combine(n, k))

	n, k = 1, 1
	fmt.Println(combine(n, k))

	n, k = 10, 5
	fmt.Println(combine(n, k))
}
