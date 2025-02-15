package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	fmt.Println(nums)
	sortColors2(nums)
	fmt.Println(nums)
}

func sortColors1(nums []int) {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	values := []int{0, 1, 2}
	index := 0
	for _, value := range values {
		count := counts[value]
		for range count {
			nums[index] = value
			index++
		}
	}
}

// based on the Dutch National Flag Algorithm
func sortColors2(nums []int) {
	low, mid := 0, 0
	high := len(nums) - 1
	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[high], nums[mid] = nums[mid], nums[high]
			high--
		}
	}
}
