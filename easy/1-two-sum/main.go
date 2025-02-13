package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	indexes := twoSum3(nums, target)
	fmt.Printf("%v\n", indexes)
}

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		m[num] = i
	}
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok && j != i {
			return []int{i, j}
		}
	}
	return []int{}
}

func twoSum3(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok && j != i {
			return []int{i, j}
		}
		m[num] = i
	}
	return []int{}
}
