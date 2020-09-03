// Design your implementation of the linked list. You can choose to use the singly linked list or the doubly linked list. A node in a singly linked list should have two attributes: val and next. val is the value of the current node, and next is a pointer/reference to the next node. If you want to use the doubly linked list, you will need one more attribute prev to indicate the previous node in the linked list. Assume all nodes in the linked list are 0-indexed.

// Implement these functions in your linked list class:

// get(index) : Get the value of the index-th node in the linked list. If the index is invalid, return -1.
// addAtHead(val) : Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
// addAtTail(val) : Append a node of value val to the last element of the linked list.
// addAtIndex(index, val) : Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
// deleteAtIndex(index) : Delete the index-th node in the linked list, if the index is valid.

// Example:

// Input:
// ["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
// [[],[1],[3],[1,2],[1],[1],[1]]
// Output:
// [null,null,null,null,2,null,3]

// Explanation:
// MyLinkedList linkedList = new MyLinkedList(); // Initialize empty LinkedList
// linkedList.addAtHead(1);
// linkedList.addAtTail(3);
// linkedList.addAtIndex(1, 2);  // linked list becomes 1->2->3
// linkedList.get(1);            // returns 2
// linkedList.deleteAtIndex(1);  // now the linked list is 1->3
// linkedList.get(1);            // returns 3

// Constraints:

// 0 <= index,val <= 1000
// Please do not use the built-in LinkedList library.
// At most 2000 calls will be made to get, addAtHead, addAtTail,  addAtIndex and deleteAtIndex.

package main

import "fmt"

type node struct {
	value int
	next  *node
}
type MyLinkedList struct {
	size int
	head *node
	tail *node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	t := &node{}
	h := &node{next: t}
	return MyLinkedList{
		head: h,
		tail: t,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || this.size <= index {
		return -1
	}
	i, cur := 0, this.head.next
	for i < index {
		i++
		cur = cur.next
	}
	return cur.value
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	new := &node{
		value: val,
	}
	new.next = this.head.next
	this.head.next = new
	this.size++
	fmt.Println(this)

}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	this.tail.value = val
	nd := &node{}
	this.tail.next = nd
	this.tail = nd
	this.size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	switch {
	case index > this.size || index < 0:
		return
	case index == 0:
		this.AddAtHead(val)
		return
	case index == this.size:
		this.AddAtTail(val)
		return
	}
	i, cur := -1, this.head
	for i+1 < index {
		i++
		cur = cur.next
	}
	nd := &node{value: val}
	nd.next = cur.next
	cur.next = nd
	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size || index < 0 {
		return
	}
	i, cur := -1, this.head
	for i+1 < index {
		i++
		cur = cur.next
	}

	cur.next = cur.next.next

	this.size--

}

func main() {

}
