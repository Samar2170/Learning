package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	// Write your code here
	pmAm, timeW := s[len(s)-2:], s[:len(s)-2]
	tArr := strings.Split(timeW, ":")
	var newHour string
	newHourInt, _ := strconv.ParseInt(tArr[0], 10, 32)
	// newMinuteInt, _ := strconv.ParseInt(tArr[1], 10, 32)
	// newSecondInt, _ := strconv.ParseInt(tArr[2], 10, 32)
	if newHourInt == 12 && pmAm == "AM" {
		newHour = "00"
	} else if pmAm == "PM" && newHourInt < 12 {
		newHour = fmt.Sprintf("%d", newHourInt+12)
	} else {
		if newHourInt >= 10 {
			newHour = fmt.Sprintf("%d", newHourInt)
		} else {
			newHour = fmt.Sprintf("0%d", newHourInt)
		}
	}
	return fmt.Sprintf("%s:%s:%s", newHour, tArr[1], tArr[2])

}

func main() {
	// all 24 hour test cases
	fmt.Println(timeConversion("12:00:00AM"))
	fmt.Println(timeConversion("12:30:00AM"))
	fmt.Println(timeConversion("01:00:00AM"))
	fmt.Println(timeConversion("01:30:00AM"))
	fmt.Println(timeConversion("02:00:00AM"))
	fmt.Println(timeConversion("02:30:00AM"))
	fmt.Println(timeConversion("06:00:00AM"))
	fmt.Println(timeConversion("06:30:00AM"))
	fmt.Println(timeConversion("09:00:00AM"))
	fmt.Println(timeConversion("09:30:00AM"))

	fmt.Println(timeConversion("10:00:00AM"))
	fmt.Println(timeConversion("10:30:00AM"))

	fmt.Println(timeConversion("12:00:00PM"))
	fmt.Println(timeConversion("12:30:00PM"))
	fmt.Println(timeConversion("06:00:00PM"))
	fmt.Println(timeConversion("06:30:00PM"))

	fmt.Println(timeConversion("11:00:00PM"))
	fmt.Println(timeConversion("11:30:00PM"))

}
