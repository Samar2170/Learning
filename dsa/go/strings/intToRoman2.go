package main

import "fmt"

var orders = []map[int]byte{
	{1: 'I', 5: 'V'},
	{1: 'X', 5: 'L'},
	{1: 'C', 5: 'D'},
	{1: 'M'},
}

func intToRoman(num int) string {
	ss := []string{}
	order := 0
	for num > 0 {
		digit := num % 10
		num = num / 10
		a := digitToRoman(digit, order)
		ss = append(ss, a)
		order++
	}
	var fs string
	for i := len(ss) - 1; i > 0; i-- {
		fs += ss[i]
	}
	return fs
}

func digitToRoman(digit int, order int) string {
	var resp string
	if digit == 4 {
		resp = string(orders[order][1])
		digit = 5
	}
	if digit == 9 {
		resp = string(orders[order][1])
		digit = 1
		order++
	}
	if digit == 5 {
		resp = string(orders[order][5])
		digit -= 5
	}
	if digit > 0 {
		for digit > 0 {
			resp += string(orders[order][1])
			digit--
		}
	}
	return resp
}

func main() {
	fmt.Println(intToRoman(1994))
	fmt.Println(intToRoman(58))
}
