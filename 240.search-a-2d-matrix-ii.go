func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 {
		return false
	}
	for i, j := len(matrix)-1, 0; i >= 0 && j < len(matrix[0]); {
		if matrix[i][j] < target {
			j++
		} else if matrix[i][j] > target {
			i--
		} else {
			return true
		}
	}
	return false
}
