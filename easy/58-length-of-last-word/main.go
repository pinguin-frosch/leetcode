package main

import (
	"fmt"
	"unicode"
)

func main() {
	word := "   fly me   to   the moon  "
	length := lengthOfLastWord1(word)
	fmt.Println(length)
}

func lengthOfLastWord1(s string) int {
	inWord := false
	length := 0
	for _, c := range s {
		if unicode.IsSpace(c) {
			inWord = false
		} else {
			if !inWord {
				length = 0
				inWord = true
			}
			length++
		}
	}
	return length
}

func lengthOfLastWord2(s string) int {
	length := 0
	end := len(s) - 1
	for unicode.IsSpace(rune(s[end])) {
		end--
	}
	for end >= 0 && !unicode.IsSpace(rune(s[end])) {
		end--
		length++
	}
	return length
}
