/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

*/

package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := new(ListNode)
	tail := head
	for l1 != nil || l2 != nil || carry != 0 {
		var l1n, l2n = 0, 0
		if l1 != nil {
			l1n, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			l2n, l2 = l2.Val, l2.Next
		}
		fmt.Println(l1n)
		fmt.Println(l2n)
		fmt.Println(carry)
		n := l1n + l2n + carry
		fmt.Println(n)
		carry = n / 10
		fmt.Println(carry)
		tail.Next = &ListNode{Val: n % 10, Next: nil}
		fmt.Println(n % 10)
		tail = tail.Next
	}
	return head.Next
}

func main() {
	addTwoNumbers()
}
