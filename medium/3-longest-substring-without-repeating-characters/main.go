package main

import "fmt"

func main() {
	s := "abcabcabcd"
	length := lengthOfLongestSubstring1(s)
	fmt.Println(length)
}

func lengthOfLongestSubstring1(s string) int {
	l, r := 0, 0
	counts := make(map[byte]int)
	longest := 0
	for r < len(s) {
		char := s[r]
		counts[char]++
		if counts[char] > 1 {
			longest = max(longest, len(counts))
			clear(counts)
			l++
			r = l
			continue
		}
		r++
	}
	return max(longest, len(counts))
}

func lengthOfLongestSubstring2(s string) int {
	l, r := 0, 0
	seen := make(map[byte]int)
	longest := 0
	for ; r < len(s); r++ {
		char := s[r]
		if index, exists := seen[char]; exists && index >= l {
			l = index + 1
		}
		seen[char] = r
		longest = max(longest, r-l+1)
	}
	return longest
}
