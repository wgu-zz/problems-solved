/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var cur *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = cur
		cur = head
		head = tmp
	}
	return cur
}
