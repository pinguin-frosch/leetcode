package main

import "fmt"

func main() {
	s := "...ww)пфикш"
	t := "пфи.w.).wкш"
	valid := isAnagram2(s, t)
	fmt.Println(valid)
}

func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countS := make(map[rune]int)
	countT := make(map[rune]int)

	for _, c := range s {
		countS[c]++
	}
	for _, c := range t {
		countT[c]++
	}

	for c, v := range countS {
		if _, ok := countT[c]; !ok {
			return false
		}
		if v != countT[c] {
			return false
		}
	}

	return true
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := make(map[rune]int)
	for _, c := range s {
		counts[c]++
	}
	for _, c := range t {
		counts[c]--
		if counts[c] < 0 {
			return false
		}
	}
	return true
}
