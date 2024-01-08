package main

import (
	"fmt"
	"strings"
)

var orders = [4]map[int]byte{
	{
		1: 'I',
		5: 'V',
	},
	{
		1: 'X',
		5: 'L',
	},
	{
		1: 'C',
		5: 'D',
	},
	{
		1: 'M',
	},
}

func intToRoman(num int) string {
	builder := &strings.Builder{}
	res := []string{}
	order := 0
	for num > 0 {
		digit := num % 10
		num = num / 10
		digitToRoman(digit, order, builder)
		order++
		res = append(res, builder.String())
		builder.Reset()
	}

	for i := len(res); i > 0; i-- {
		builder.WriteString(res[i])
	}
	return builder.String()
}

func digitToRoman(digit int, order int, builder *strings.Builder) {
	if digit == 4 {
		builder.WriteByte(orders[order][1])
		digit = 5
	}
	if digit == 9 {
		builder.WriteByte(orders[order][1])
		digit = 1
		order = order + 1
	}
	// if digit 4 or 9 we add I in the beginning, in case of 9
	//  inc order by 1 to get X
	if digit <= 5 {
		builder.WriteByte(orders[order][5])
		digit -= 5
	}
	for digit > 0 {
		digit--
		builder.WriteByte(orders[order][1])
	}

}

func main() {
	nums := []int{3, 4, 9, 58, 1994}
	// for _, num := range nums {
	// 	fmt.Println(intToRoman(num))
	// }
	fmt.Println(intToRoman(nums[3]))
}
