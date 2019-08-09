/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 *
 * https://leetcode.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (30.95%)
 * Total Accepted:    252.7K
 * Total Submissions: 814.7K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * Given a matrix of m x n elements (m rows, n columns), return all elements of
 * the matrix in spiral order.
 *
 * Example 1:
 *
 *
 * Input:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 * Output: [1,2,3,6,9,8,7,4,5]
 *
 *
 * Example 2:
 *
 * Input:
 * [
 * ⁠ [1, 2, 3, 4],
 * ⁠ [5, 6, 7, 8],
 * ⁠ [9,10,11,12]
 * ]
 * Output: [1,2,3,4,8,12,11,10,9,5,6,7]
 *
 */

func spiralOrder(matrix [][]int) []int {
	if len(matrix) < 1 {
		return []int{}
	}
	res, dir := make([]int, len(matrix)*len(matrix[0])), 0
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	i, j, x := 0, 0, 0
	for i >= top && i <= bottom && j >= left && j <= right {
		res[x] = matrix[i][j]
		switch dir {
		case 0:
			if j == right {
				i++
				dir++
				top++
			} else {
				j++
			}
		case 1:
			if i == bottom {
				j--
				dir++
				right--
			} else {
				i++
			}
		case 2:
			if j == left {
				i--
				dir++
				bottom--
			} else {
				j--
			}
		case 3:
			if i == top {
				j++
				dir++
				left++
			} else {
				i--
			}
		default:
		}
		dir %= 4
		x++
	}
	return res
}
