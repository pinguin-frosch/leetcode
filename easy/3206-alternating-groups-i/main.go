package main

import "fmt"

func main() {
	colors := []int{0,1,0,0,1}
	number := numberOfAlternatingGroups(colors)
	fmt.Println(number)
}

func numberOfAlternatingGroups(colors []int) int {
	number := 0
	n := len(colors)
	for i := 0; i < n; i++ {
		prev := colors[(i-1+n) % n]
		curr := colors[i]
		next := colors[(i+1)%n]
		if prev == next && prev != curr {
			number++
		}
	}
	return number
}
