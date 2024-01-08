package main

import (
	"fmt"
	"math"
)

func primeNumbers(start int, end int) []int {
	if start < 2 || end < 2 {
		fmt.Println("Numbers must be greater than 2.")
		return []int{}
	}
	ans := []int{}
	for start <= end {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(start))); i++ {
			if start%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			ans = append(ans, start)
		}
		start++
	}
	return ans
}
func main() {
	fmt.Println(primeNumbers(5, 19))
	fmt.Println(primeNumbers(0, 2))
	fmt.Println(primeNumbers(13, 100))
}
