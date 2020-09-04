package main

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev //For the first time this will be set to nil which we make the last node to be empty/nil
		prev = head      //to add this to the end
		head = temp
	}
	return prev
}
