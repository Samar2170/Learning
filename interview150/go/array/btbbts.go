package main

import "fmt"

func maxProfit(prices []int) int {
	minPrice := prices[0]
	mProfit := 0
	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		}
		profit := p - minPrice
		if profit > mProfit {
			mProfit = profit
		}
	}
	return mProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}
