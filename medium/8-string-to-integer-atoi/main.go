package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)

func main() {
	string := "42"
	int := myAtoi2(string)
	fmt.Println(int)
}

func myAtoi1(s string) int {
	sign := 1
	digits := ""
	number := 0

	// ignore leading whitespace and check sign
	digitsIndex := 0
	for i, c := range s {
		if c == '-' {
			sign = -1
			digitsIndex = i + 1
			break
		}
		if c == '+' {
			digitsIndex = i + 1
			break
		}
		if unicode.IsDigit(c) {
			digitsIndex = i
			break
		}
		if unicode.IsSpace(c) {
			continue
		} else {
			break
		}
	}

	// invalid case: ends with -
	if digitsIndex >= len(s) {
		return 0
	}

	// invalid case: - followed by no digits
	if !unicode.IsDigit(rune(s[digitsIndex])) {
		return 0
	}

	// record all the digits
	for i := digitsIndex; i < len(s); i++ {
		digit := rune(s[i])
		if !unicode.IsDigit(digit) {
			break
		}
		digits += string(digit)
	}

	// remove leading 0
	nonLeadingZeroDigits := ""
	inLeadingZero := true
	for _, c := range digits {
		if c != '0' {
			inLeadingZero = false
		}
		if !inLeadingZero {
			nonLeadingZeroDigits += string(c)
		}
	}

	digits = nonLeadingZeroDigits

	if digits == "" {
		return 0
	}
	// no int32 value can have 10 or more digits
	if len(digits) > 10 {
		if sign == -1 {
			return -2147483648
		} else {
			return 2147483647
		}
	}

	number, err := strconv.Atoi(digits)
	if err != nil {
		log.Fatal(err)
	}
	number *= sign

	if number < 0 && number < -2147483648 {
		number = -2147483648
		fmt.Println(1)
	} else if number > 2147483647 {
		number = 2147483647
		fmt.Println(2)
	}

	return number
}

func myAtoi2(s string) int {
	// ignore whitespace
	index := 0
	for index < len(s) && unicode.IsSpace(rune(s[index])) {
		index++
	}

	// check for sign
	sign, number := 1, 0
	if index >= len(s) {
		return 0
	}
	if string(s[index]) == "-" {
		sign = -1
		index++
	} else if string(s[index]) == "+" {
		sign = 1
		index++
	}

	// values to convert into an integer
	values := map[rune]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
		'5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
	}
	const MAX_INT = 2147483647

	// construct the final number
	for index < len(s) && unicode.IsDigit(rune(s[index])) {
		digit := values[rune(s[index])]

		// check boundaries for int32
		if number > MAX_INT/10 || (number == MAX_INT/10 && digit > 7) {
			if sign == 1 {
				return MAX_INT
			} else {
				return -(MAX_INT + 1)
			}
		}
		number = number*10 + digit
		index++
	}

	return sign * number
}
