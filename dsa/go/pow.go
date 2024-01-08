package main

import "fmt"

func myPow(x float64, n int) float64 {
	var ans float64 = 1
	if n > 0 {
		i := 0
		for i < n {
			ans = ans * x
			i++
		}

	} else if n < 0 {
		var dem float64
		dem = 1
		i := 0
		for i < -n {
			dem = dem * x
			i++
		}
		ans = 1 / dem
	} else {
		ans = 1
	}
	return ans
}

func main() {
	fmt.Println(myPow(2, 3))
	fmt.Println(myPow(2, -2))
	fmt.Println(myPow(2, 0))
	fmt.Println(myPow(0.00001, 2147483647))

	// print(sol.myPow(2,-2))
	// print(sol.myPow(2,0))
	// print(sol.myPow(0.00001,2147483647))

}
