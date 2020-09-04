// Merge two sorted linked lists and return it as a new sorted list. The new list should be made by splicing together the nodes of the first two lists.

// Example:

// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4
package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var ln *ListNode
	out := ln
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			ln.Next = l1
			l1 = l1.Next
			ln = ln.Next
		} else {
			ln.Next = l2
			l2 = l2.Next
			ln = ln.Next
		}
	}
	if l1 != nil {
		ln.Next = l1
	} else if l2 != nil {
		ln.Next = l2
	}

	return out.Next
}
