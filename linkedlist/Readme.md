## Linked List

Linked List is linear data structure liked Array. Each element in the linked list is actually a separate object while all the objects are linked together by the reference field in each element.

There are two types of linked list: singly linked list and doubly linked list. 

### Singly Linked list
#### Introduction
Singly list contains not only the value but also a reference field to the next node.Singly linked list organizes all the node in sequence.

Advantage:
Unlike the array, we are not able to access a random element in a singly-linked list in constant time. If we want to get the ith element, we have to traverse from the head node one by one. It takes us O(N) time on average to visit an element by index, where N is the length of the linked list.

#### Two-Pointer in Linked List
Imagine there are two runners with different speed. If they are running on a straight path, the fast runner will first arrive at the destination. However, if they are running on a circular track, the fast runner will catch up with the slow runner if they keep running.

That's exactly what we will come across using two pointers with different speed in a linked list:

If there is no cycle, the fast pointer will stop at the end of the linked list.
If there is a cycle, the fast pointer will eventually meet with the slow pointer.