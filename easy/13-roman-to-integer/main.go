package main

import "fmt"

func main() {
	roman := "MCMXCIV"
	number := romanToInt1(roman)
	fmt.Println(number)
}

func romanToInt1(s string) int {
	values := map[rune]int{
		0:   0,
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	number := 0
	for i := 0; i < len(s); i++ {
		current := rune(s[i])
		next := rune(0)
		if i != len(s)-1 {
			next = rune(s[i+1])
		}
		sign := 1
		if values[current] < values[next] {
			sign = -1
		}
		number += sign * values[current]
	}
	return number
}
