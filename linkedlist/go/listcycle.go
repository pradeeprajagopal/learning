package main

// Given a linked list, determine if it has a cycle in it.

// To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where the tail connects to. If pos == -1, then there is no cycle in the linked list.

// Follow up:

// Can you solve it using O(1) (i.e. constant) memory?

// Example 1:

// Input: head = [3,2,0,-4], pos = 1
// Output: true
// Explanation: There is a cycle in the linked list, where tail connects to the second node.
// Example 2:

// Input: head = [1,2], pos = 0
// Output: true
// Explanation: There is a cycle in the linked list, where tail connects to the first node.
// Example 3:

// Input: head = [1], pos = -1
// Output: false
// Explanation: There is no cycle in the linked list.

// Constraints:

// The number of the nodes in the list is in the range [0, 104].
// -105 <= Node.val <= 105
// pos is -1 or a valid index in the linked-list.

func detectCycle(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return nil
	}
	s, f := head.Next, head.Next.Next
	for f != nil && f.Next != nil && s != f {
		s, f = s.Next, f.Next.Next

	}
	if s != f {
		return nil
	}
	for s != head {
		s, head = s.Next, head.Next
	}

	return slow
}
