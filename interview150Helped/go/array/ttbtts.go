package main

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	minPrice := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}
	return profit
}
func main() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices1))
	prices2 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices2))

	prices3 := []int{1, 2}
	fmt.Println(maxProfit(prices3))

}
