package main

func climbStairs(n int) int {
	next, secondNext := 0, 1
	i := 0
	for i < n {
		next, secondNext = secondNext, next+secondNext
		i++
	}
	return secondNext
}
