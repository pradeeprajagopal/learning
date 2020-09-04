// Given a singly linked list, determine if it is a palindrome.

// Example 1:

// Input: 1->2
// Output: false
// Example 2:

// Input: 1->2->2->1
// Output: true
// Follow up:
// Could you do it in O(n) time and O(1) space?

package main

func isPalindrome(head *ListNode) bool {
	nums := make([]int, 0, 64)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	i := 0
	count := len(nums) - 1
	for i < count {
		if nums[i] != nums[count] {
			return false
		}
		i++
		count--
	}
	return true

}
