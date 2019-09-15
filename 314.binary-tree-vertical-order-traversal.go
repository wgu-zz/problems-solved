/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	m := make(map[int][]int)
	q := []*Node{&Node{root, 0}}
	min, max := 0, 0
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		m[n.index] = append(m[n.index], n.tn.Val)
		if n.tn.Left != nil {
			leftCol := n.index - 1
			if leftCol < min {
				min = leftCol
			}
			q = append(q, &Node{n.tn.Left, leftCol})
		}
		if n.tn.Right != nil {
			rightCol := n.index + 1
			if rightCol > max {
				max = rightCol
			}
			q = append(q, &Node{n.tn.Right, rightCol})
		}
	}
	res := [][]int{}
	for i := min; i <= max; i++ {
		res = append(res, m[i])
	}
	return res
}

type Node struct {
	tn    *TreeNode
	index int
}
