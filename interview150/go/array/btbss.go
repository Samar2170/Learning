package main

// self
func maxProfit(prices []int) int {
	mp := 0
	minPrice := prices[0]
	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		}
		profit := p - minPrice
		if profit > mp {
			mp = profit
		}
	}
	return mp
}

func main() {
	var prices []int
	prices = []int{7, 1, 5, 3, 6, 4}
	println(maxProfit(prices))

	prices = []int{7, 6, 4, 3, 1}
	println(maxProfit(prices))
}
