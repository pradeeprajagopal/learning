package main

/*
369. Plus One Linked List

Given a non-negative integer represented as non-empty a singly linked list of digits, plus one to the integer.

You may assume the integer do not contain any leading zero, except the number 0 itself.

The digits are stored such that the most significant digit is at the head of the list.

Example :

Input: [1,2,3]
Output: [1,2,4]
*/
func plusOne(head *ListNode) *ListNode {
	next := carry(head)
	if next {
		newHead := ListNode{1, head}
		return &newHead
	} else {
		return head
	}
}

func carry(head *ListNode) bool {
	if head == nil {
		return true
	}
	next := carry(head.Next)
	curVal := head.Val
	if next {
		curVal += 1
	}
	if curVal >= 10 {
		head.Val = curVal % 10
		return true
	} else {
		head.Val = curVal
		return false
	}
}
