package main

import (
	"fmt"
	"math"
)

func main() {
	number := -123
	reversed := reverse(number)
	fmt.Println(reversed)
}

func reverse(x int) int {
	// save sign for later
	sign := 1
	if x < 0 {
		sign = -1
	}

	// make the number positive
	if sign == -1 {
		x *= sign
	}

	// reverse the number
	reversed := 0
	for x > 0 {
		digit := x % 10
		if reversed > math.MaxInt32/10 || (reversed == math.MaxInt32/10 && digit > 7) {
			return 0
		}
		reversed = reversed*10 + digit
		x /= 10
	}

	return sign * reversed
}
