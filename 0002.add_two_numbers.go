package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result, carry := l1, 0
	var cur *ListNode
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			cur = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			cur = l2
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if sum/10 == 0 {
			cur.Val = sum
			carry = 0
		} else {
			cur.Val = sum % 10
			carry = 1
		}
	}
	if carry == 1 {
		cur.Next = &ListNode{1, nil}
	}
	return result
}
