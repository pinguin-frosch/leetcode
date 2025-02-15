package main

import (
	"fmt"
	"pinguinfrosch.com/utils"
)

func main() {
	s := "([[}]])"
	valid := isValid(s)
	fmt.Println(valid)

}

func isValid(s string) bool {
	stack := utils.NewStack[rune]()
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack.Push(c)
		} else {
			if stack.Len() == 0 {
				return false
			}
			obj, _ := stack.Pop()
			if c == ']' && obj != '[' {
				return false
			}
			if c == ')' && obj != '(' {
				return false
			}
			if c == '}' && obj != '{' {
				return false
			}
		}
	}
	return stack.Len() == 0
}
