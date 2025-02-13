package main

import (
	"fmt"
)

func main() {
	l1 := digitsToListNode([]int{2, 4, 9})
	l2 := digitsToListNode([]int{5, 6, 4, 9})
	result := addTwoNumbers2(l1, l2)
	digitsResult := listNodeToDigits(result)
	fmt.Println(digitsResult)
}

// creates a ListNode more easily by just using the digits
func digitsToListNode(digits []int) *ListNode {
	nodes := make([]ListNode, 0, len(digits))
	for _, d := range digits {
		n := ListNode{Val: d}
		nodes = append(nodes, n)
	}

	// add the next pointer for each node
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = &nodes[i+1]
	}

	return &nodes[0]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	digits1 := listNodeToDigits(l1)
	digits2 := listNodeToDigits(l2)

	// create new slice for the sum of each pair of digits, value if 0 if not
	// present
	maxLength := max(len(digits1), len(digits2))
	digitsResult := make([]int, maxLength)
	carry := 0
	for i := 0; i < maxLength; i++ {
		d1, d2 := 0, 0
		if i < len(digits1) {
			d1 = digits1[i]
		}
		if i < len(digits2) {
			d2 = digits2[i]
		}
		sum := d1 + d2 + carry
		carry = sum / 10
		digitsResult[i] = sum % 10
	}

	// add a new digit if there's carry left
	if carry > 0 {
		digitsResult = append(digitsResult, carry)
	}

	// form the respective nodes
	nodes := make([]ListNode, 0, len(digitsResult))
	for _, d := range digitsResult {
		n := ListNode{Val: d}
		nodes = append(nodes, n)
	}

	// add the next pointer for each node
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = &nodes[i+1]
	}

	return &nodes[0]
}

// retrieves the digits from the ListNode
func listNodeToDigits(l *ListNode) []int {
	digits := make([]int, 0)
	// add the digits of the first nodes
	for l.Next != nil {
		digits = append(digits, l.Val)
		l = l.Next
	}
	// add the digit of the last node
	digits = append(digits, l.Val)

	return digits
}

// i didn't write this myself. It's essentially what i did but everything on
// one pass, it's much much simpler
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}
	return result.Next
}
