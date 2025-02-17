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
			// TODO: only remove the repeated element
			clear(counts)
			l++
			r = l
			continue
		}
		r++
	}
	return max(longest, len(counts))
}
