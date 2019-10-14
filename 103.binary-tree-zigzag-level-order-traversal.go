/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	res := [][]int{}
	index := 0
	for len(q) > 0 {
		l := len(q)
		row := []int{}
		if index == 0 {
			for i := 0; i < l; i++ {
				row = append(row, q[i].Val)
			}
		} else {
			for i := l - 1; i >= 0; i-- {
				row = append(row, q[i].Val)
			}
		}
		res = append(res, row)
		index = (index + 1) % 2
		for i := 0; i < l; i++ {
			n := q[0]
			q = q[1:]
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
	}
	return res
}
