package main

func removeElements(head *ListNode, val int) *ListNode {
	prev := &ListNode{Next: head}
	temp := prev
	for temp.Next != nil {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}

	return prev.Next
}
