/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	preHead := &ListNode{Next: head}
	p := 1
	subHead := preHead
	for p != m && head != nil {
		subHead = head
		head = head.Next
		p++
	}
	subTail := head
	var prev, next *ListNode
	for p != n+1 && head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
		p++
	}
	subHead.Next = prev
	if subTail != nil {
		subTail.Next = head
	}
	return preHead.Next
}
