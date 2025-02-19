package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9}
	k := 3
	duplicate := containsNearbyDuplicate(nums, k)
	fmt.Println(duplicate)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int)
	for i, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = i
		if len(seen) > k {
			delete(seen, nums[i-k])
		}
	}
	return false
}
