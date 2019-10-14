/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flattenTail(root)
}

func flattenTail(node *TreeNode) *TreeNode {
	if node.Left == nil && node.Right == nil {
		return node
	}
	if node.Left != nil {
		right := node.Right
		node.Right = node.Left
		node.Left = nil
		tail := flattenTail(node.Right)
		tail.Right = right
		if right != nil {
			return flattenTail(right)
		}
		return tail
	} else {
		return flattenTail(node.Right)
	}
}
