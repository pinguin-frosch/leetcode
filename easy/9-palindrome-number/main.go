package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := 121
	result := isPalindrome1(number)
	fmt.Println(result)
}

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	string := strconv.Itoa(x)
	for i := 0; i < len(string)/2; i++ {
		a, b := string[i], string[len(string)-1-i]
		if a != b {
			return false
		}
	}
	return true
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	digits := make([]int, 0)
	for x != 0 {
		digit := x % 10
		x = x / 10
		digits = append(digits, digit)
	}
	for i := 0; i < len(digits)/2; i++ {
		a, b := digits[i], digits[len(digits)-1-i]
		if a != b {
			return false
		}
	}
	return true
}

func isPalindrome3(x int) bool {
	target := x
	if x < 0 {
		return false
	}
	sum := 0
	for x != 0 {
		digit := x % 10
		x = x / 10
		sum = sum*10 + digit
	}
	return sum == target
}
