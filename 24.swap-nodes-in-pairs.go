/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	res := head.Next
	for head != nil && head.Next != nil {
		n1, n2 := head, head.Next
		head = head.Next.Next
		if prev != nil {
			prev.Next = n2
		}
		n2.Next = n1
		n1.Next = head
		prev = n1
	}
	return res
}
