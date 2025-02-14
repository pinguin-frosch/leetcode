package main

import (
	"fmt"
	"math"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	longest := longestCommonPrefix(strs)
	fmt.Println(longest)
}

func longestCommonPrefix(strs []string) string {
	prefix := ""
	minimumLength := math.MaxInt
	for _, str := range strs {
		if len(str) < minimumLength {
			minimumLength = len(str)
		}
	}
	for i := 0; i < minimumLength; i++ {
		common := ""
		for j, str := range strs {
			char := string(str[i])
			if j == 0 {
				common = char
			}
			if char != common {
				return prefix
			}
		}
		prefix += common
	}
	return prefix
}
