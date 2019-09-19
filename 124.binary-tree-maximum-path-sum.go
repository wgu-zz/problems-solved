/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */import "math"

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	dfs(root, &res)
	return res
}

func dfs(node *TreeNode, res *int) int {
	if node == nil {
		return 0
	}
	value := node.Val
	leftGain := max(0, dfs(node.Left, res))
	value += leftGain
	rightGain := max(0, dfs(node.Right, res))
	value += rightGain
	*res = max(*res, value)
	return node.Val + max(leftGain, rightGain)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
