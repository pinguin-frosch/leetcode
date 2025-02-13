package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	hasDuplicates := containsDuplicate(nums)
	fmt.Println(hasDuplicates)
}

func containsDuplicate(nums []int) bool {
	found := make(map[int]bool)
	for _, n := range nums {
		if v := found[n]; v {
			return true
		}
		found[n] = true
	}
	return false
}
