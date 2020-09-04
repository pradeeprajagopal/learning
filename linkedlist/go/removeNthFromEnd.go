// Given a linked list, remove the n-th node from the end of list and return its head.

// Example:

// Given linked list: 1->2->3->4->5, and n = 2.

// After removing the second node from the end, the linked list becomes 1->2->3->5.
// Note:

// Given n will always be valid.

// Follow up:

// Could you do this in one pass?
package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ptr := head
	empty := &ListNode{
		Next: &ListNode
	}
	var ctr,length int
	for ptr != nil {
		ptr = ptr.Next
		ctr++
	}
	length = ctr - n
	prev := &ListNode{Next:empty}
	for length < 0 {
		prev = ptr
		ptr = ptr.Next
		length--
	}
	if prev.Next == nil {
		return head.Next
	}
	prev.Next = ptr.Next
	ptr.Next = &ListNode{}
	return head

}
