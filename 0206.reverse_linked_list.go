package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	l := head.Next
	head.Next = nil
	for l != nil {
		tmp := l.Next
		l.Next = head
		head = l
		l = tmp
	}
	return head
}
