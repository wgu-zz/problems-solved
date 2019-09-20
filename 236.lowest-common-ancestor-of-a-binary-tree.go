/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

var res *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	find(root, p, q)
	return res
}

func find(node, p, q *TreeNode) int {
	if node == nil {
		return 0
	}
	leftSeen := find(node.Left, p, q)
	rightSeen := find(node.Right, p, q)
	thisSeen := 0
	if node == p || node == q {
		thisSeen = 1
	}
	if leftSeen+rightSeen+thisSeen > 1 {
		res = node
	}
	return max(thisSeen, max(leftSeen, rightSeen))
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
