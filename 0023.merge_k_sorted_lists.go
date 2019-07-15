package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	listLength := len(lists)
	if listLength <= 0 {
		return nil
	}
	for interval := 1; interval < listLength; interval *= 2 {
		for i := 0; i < listLength-interval; i += interval * 2 {
			lists[i] = merge2Lists(lists[i], lists[i+interval])
		}
	}
	return lists[0]
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	point := head
	for l1 != nil && l2 != nil {
		smaller := l1
		if l2.Val < l1.Val {
			smaller = l2
			l2 = l2.Next
		} else {
			l1 = l1.Next
		}
		point.Next = smaller
		point = point.Next
	}
	if l1 != nil {
		point.Next = l1
	} else {
		point.Next = l2
	}
	return head.Next
}
